package main

import (
	"embed"
	"encoding/json"
	"github.com/corlys/corlys.dev/model"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed templates data assets
var files embed.FS

// card is the shape templates/projects.html ranges over.
type card struct {
	Title, Description, Link, Image string
	Internal                        bool
}

var (
	pages    = load()
	projects = loadProjects()
)

func load() map[string]*template.Template {
	pages := map[string]*template.Template{}
	for _, p := range []string{"home", "project", "projects"} {
		pages[p] = template.Must(template.ParseFS(files, "templates/layout.html", "templates/"+p+".html"))
	}
	return pages
}

func loadProjects() []model.Project {
	b, err := files.ReadFile("data/projects.json")
	if err != nil {
		log.Fatal(err)
	}
	var ps []model.Project
	if err := json.Unmarshal(b, &ps); err != nil {
		log.Fatal(err)
	}
	return ps
}

func render(w http.ResponseWriter, page string, data any) {
	if err := pages[page].ExecuteTemplate(w, "layout", data); err != nil {
		log.Println(err)
		http.Error(w, "render error", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "home", nil)
}

func categoryHandler(category, heading, accent string) http.HandlerFunc {
	var cards []card
	for _, p := range projects {
		if p.Category != category {
			continue
		}
		link := p.Demo
		if link == "" {
			link = p.Repo
		}
		cards = append(cards, card{Title: p.Title, Description: p.Summary, Link: link, Image: p.Image})
	}
	data := map[string]any{
		"Heading":   heading,
		"Back":      "/",
		"BackLabel": "Back",
		"Accent":    accent,
		"Projects":  cards,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "projects", data)
	}
}

func projectHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()

	assets, _ := fs.Sub(files, "assets")

	mux.HandleFunc("GET /{$}", homeHandler)
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServerFS(assets)))
	mux.HandleFunc("GET /frontend", categoryHandler("frontend", "Frontend Projects", "green"))
	mux.HandleFunc("GET /backend", categoryHandler("backend", "Backend Projects", "blue"))
	mux.HandleFunc("GET /{category}/{id}", projectHandler)

	log.Fatal(http.ListenAndServe(":8888", mux))
}

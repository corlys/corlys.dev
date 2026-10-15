package model

type Project struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Summary     string   `json:"summary"`
	Description []string `json:"description"`
	Stack       []string `json:"stack"`
	Role        string   `json:"role"`
	Year        int      `json:"year"`
	Repo        string   `json:"repo"`
	Demo        string   `json:"demo"`
	Image       string   `json:"image"`
}

type Module struct {
	Title       string   `json:"title"`
	SubHeading  string   `json:"subHeading"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	KeyFeatures []string `json:"keyFeatures"`
}

type Detail struct {
	Modules []Module `json:"modules"`
	Type    string   `json:"type"`
	Name    string   `json:"name"`
}

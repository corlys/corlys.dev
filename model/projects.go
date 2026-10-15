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

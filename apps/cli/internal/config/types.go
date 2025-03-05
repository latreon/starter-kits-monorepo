package config

type StarterKit struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Structure   []string `json:"structure"`
	Pros        []string `json:"pros"`
	Cons        []string `json:"cons"`
}

type StarterKitsData struct {
	StarterKits []StarterKit `json:"starterKits"`
} 
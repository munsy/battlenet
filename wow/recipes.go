package wow

// Recipe represents a World of Warcraft recipe.
type Recipe struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Icon       string `json:"icon"`
}

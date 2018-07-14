package d3

// Get artisan
type ItemProduced struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}

type ReagentItem struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}
type Reagent struct {
	Quantity int         `json:"quantity"`
	Item     ReagentItem `json:"item"`
}

// Recipe is also returned from an endpoint.
type Recipe struct {
	ID           string       `json:"id"`
	Slug         string       `json:"slug"`
	Name         string       `json:"name"`
	Cost         int          `json:"cost"`
	Reagents     []Reagent    `json:"reagents"`
	ItemProduced ItemProduced `json:"itemProduced"`
}

type Tiers struct {
	Tier           int      `json:"tier"`
	TrainedRecipes []Recipe `json:"trainedRecipes"`
	TaughtRecipes  []Recipe `json:"taughtRecipes"`
}

type Training struct {
	Tiers []Tiers `json:"tiers"`
}

// Artisan holds all of the above structs when unmarshalled.
type Artisan struct {
	Slug     string   `json:"slug"`
	Name     string   `json:"name"`
	Portrait string   `json:"portrait"`
	Training Training `json:"training"`
}

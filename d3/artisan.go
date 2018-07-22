package d3

// ItemProduced represents a Diablo III item produced from a recipe.
type ItemProduced struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}

// ReagentItem represents a Diablo III reagent item.
type ReagentItem struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}

// Reagent represents a Diablo III reagent.
type Reagent struct {
	Quantity int         `json:"quantity"`
	Item     ReagentItem `json:"item"`
}

// Recipe represents a Diablo III recipe.
type Recipe struct {
	ID           string       `json:"id"`
	Slug         string       `json:"slug"`
	Name         string       `json:"name"`
	Cost         int          `json:"cost"`
	Reagents     []Reagent    `json:"reagents"`
	ItemProduced ItemProduced `json:"itemProduced"`
}

// Tiers represents Diablo III tiers.
type Tiers struct {
	Tier           int      `json:"tier"`
	TrainedRecipes []Recipe `json:"trainedRecipes"`
	TaughtRecipes  []Recipe `json:"taughtRecipes"`
}

// Training represents Diablo III training.
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

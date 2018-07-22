package d3

// ItemTypeIndex holds all of the possible Diablo III item types.
type ItemTypeIndex []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// ItemType contains more information about each Diablo III item type.
type ItemType []struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}

package d3

// Item type index
type ItemTypeIndex []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type ItemType []struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}

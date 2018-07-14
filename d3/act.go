package d3

type Quest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Act struct {
	Slug   string  `json:"slug"`
	Number int     `json:"number"`
	Name   string  `json:"name"`
	Quests []Quest `json:"quests"`
}

type ActIndex struct {
	Acts []Act `json:"acts"`
}

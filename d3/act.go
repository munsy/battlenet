package d3

// Quest represents a Diablo III quest.
type Quest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Act represents a Diablo III act.
type Act struct {
	Slug   string  `json:"slug"`
	Number int     `json:"number"`
	Name   string  `json:"name"`
	Quests []Quest `json:"quests"`
}

// ActIndex represents a Diablo III act index.
type ActIndex struct {
	Acts []Act `json:"acts"`
}

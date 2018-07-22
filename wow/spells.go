package wow

// Spell is a representation of a World of Warcraft spell.
type Spell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Range       string `json:"range"`
	PowerCost   string `json:"powerCost"`
	CastTime    string `json:"castTime"`
}

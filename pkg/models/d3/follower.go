package d3

// Follower represents a single Diablo III follower for a character.
type Follower struct {
	Slug     string  `json:"slug"`
	Name     string  `json:"name"`
	RealName string  `json:"realName"`
	Portrait string  `json:"portrait"`
	Skills   []Skill `json:"skills"`
}

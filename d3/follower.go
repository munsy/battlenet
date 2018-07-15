package d3

type Follower struct {
	Slug     string  `json:"slug"`
	Name     string  `json:"name"`
	RealName string  `json:"realName"`
	Portrait string  `json:"portrait"`
	Skills   []Skill `json:"skills"`
}

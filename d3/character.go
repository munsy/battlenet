package d3

type Skill struct {
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	Level           int    `json:"level"`
	TooltipURL      string `json:"tooltipUrl"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"descriptionHtml"`
	FlavorText      string `json:"flavorText"`
}

type Skills struct {
	Active  []Skill `json:"active"`
	Passive []Skill `json:"passive"`
}

type SkillCategory struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Get character class
type CharacterClass struct {
	Slug            string          `json:"slug"`
	Name            string          `json:"name"`
	MaleName        string          `json:"maleName"`
	FemaleName      string          `json:"femaleName"`
	Icon            string          `json:"icon"`
	SkillCategories []SkillCategory `json:"skillCategories"`
	Skills          Skills          `json:"skills"`
}

// Get API Skill
type CharacterAPISkill struct {
	Skill Skill  `json:"skill"`
	Runes []Rune `json:"runes"`
}

package d3

// Skill represents a Diablo III skill.
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

// Skills represents Diablo III skills.
type Skills struct {
	Active  []Skill `json:"active"`
	Passive []Skill `json:"passive"`
}

// SkillCategory represents a Diablo III skill category.
type SkillCategory struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// CharacterClass represents a Diablo III character class.
type CharacterClass struct {
	Slug            string          `json:"slug"`
	Name            string          `json:"name"`
	MaleName        string          `json:"maleName"`
	FemaleName      string          `json:"femaleName"`
	Icon            string          `json:"icon"`
	SkillCategories []SkillCategory `json:"skillCategories"`
	Skills          Skills          `json:"skills"`
}

// CharacterAPISkill represents a Diablo III character skill.
type CharacterAPISkill struct {
	Skill Skill   `json:"skill"`
	Runes []Skill `json:"runes"`
}

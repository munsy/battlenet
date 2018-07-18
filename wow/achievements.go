package wow

type Achievement struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Points      int          `json:"points"`
	Description string       `json:"description"`
	Reward      string       `json:"reward"`
	RewardItems []RewardItem `json:"rewardItems"`
	Icon        string       `json:"icon"`
	Criteria    []Criteria   `json:"criteria"`
	AccountWide bool         `json:"accountWide"`
	FactionID   int          `json:"factionId"`
}

type TooltipParams struct {
	TimewalkerLevel int `json:"timewalkerLevel"`
}

type Appearance struct {
}

type RewardItem struct {
	ID                   int           `json:"id"`
	Name                 string        `json:"name"`
	Icon                 string        `json:"icon"`
	Quality              int           `json:"quality"`
	ItemLevel            int           `json:"itemLevel"`
	TooltipParams        TooltipParams `json:"tooltipParams"`
	Stats                []interface{} `json:"stats"`
	Armor                int           `json:"armor"`
	Context              string        `json:"context"`
	BonusLists           []interface{} `json:"bonusLists"`
	ArtifactID           int           `json:"artifactId"`
	DisplayInfoID        int           `json:"displayInfoId"`
	ArtifactAppearanceID int           `json:"artifactAppearanceId"`
	ArtifactTraits       []interface{} `json:"artifactTraits"`
	Relics               []interface{} `json:"relics"`
	Appearance           Appearance    `json:"appearance"`
}

type Criteria struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	OrderIndex  int    `json:"orderIndex"`
	Max         int    `json:"max"`
}

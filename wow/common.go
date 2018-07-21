package wow

/*
	Types common to multiple other types go in here.
*/

type CharacterAchievement struct {
	ID          int                            `json:"id"`
	Title       string                         `json:"title"`
	Points      int                            `json:"points"`
	Description string                         `json:"description"`
	Reward      string                         `json:"reward"`
	RewardItems []interface{}                  `json:"rewardItems"`
	Icon        string                         `json:"icon"`
	Criteria    []CharacterAchievementCriteria `json:"criteria"`
	AccountWide bool                           `json:"accountWide"`
	FactionID   int                            `json:"factionId"`
}

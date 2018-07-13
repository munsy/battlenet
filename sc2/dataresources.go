package sc2

type RewardData struct {
	Title         string `json:"title"`
	ID            int64  `json:"id"`
	Icon          Icon   `json:"icon"`
	AchievementID int    `json:"achievementId"`
}

type RewardsData struct {
	Portraits     []RewardData `json:"portraits"`
	TerranDecals  []RewardData `json:"terranDecals"`
	ZergDecals    []RewardData `json:"zergDecals"`
	ProtossDecals []RewardData `json:"protossDecals"`
	Skins         []RewardData `json:"skins"`
	Animations    []RewardData `json:"animations"`
}

type AchievementData struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	AchievementID int64  `json:"achievementId"`
	CategoryID    int    `json:"categoryId"`
	Points        int    `json:"points"`
	Icon          Icon   `json:"icon"`
}

type Child struct {
	Title                 string `json:"title"`
	CategoryID            int    `json:"categoryId"`
	FeaturedAchievementID int    `json:"featuredAchievementId"`
}

type Category struct {
	Title                 string  `json:"title"`
	CategoryID            int     `json:"categoryId"`
	FeaturedAchievementID int64   `json:"featuredAchievementId"`
	Children              []Child `json:"children,omitempty"`
}

type AchievementsData struct {
	Achievements []AchievementData `json:"achievements"`
	Categories   []Category        `json:"categories"`
}

package wow

type Battlegroup struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type BattleGroupsData struct {
	Battlegroups []Battlegroup `json:"battlegroups"`
}

type CharacterRace struct {
	ID   int    `json:"id"`
	Mask int    `json:"mask"`
	Side string `json:"side"`
	Name string `json:"name"`
}

type CharacterRacesData struct {
	Races []CharacterRace `json:"races"`
}

type CharacterClass struct {
	ID        int    `json:"id"`
	Mask      int    `json:"mask"`
	PowerType string `json:"powerType"`
	Name      string `json:"name"`
}

type CharacterClassesData struct {
	Classes []CharacterClass `json:"classes"`
}

type AchievementCriteria struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	OrderIndex  int    `json:"orderIndex"`
	Max         int    `json:"max"`
}

type CharacterAchievementCategory struct {
	ID           int                    `json:"id"`
	Achievements []CharacterAchievement `json:"achievements"`
	Name         string                 `json:"name"`
}

type CharacterAchievementData struct {
	ID           int                            `json:"id"`
	Achievements []CharacterAchievement         `json:"achievements,omitempty"`
	Name         string                         `json:"name"`
	Categories   []CharacterAchievementCategory `json:"categories,omitempty"`
}

type CharacterAchievementsData struct {
	Achievements []CharacterAchievementData `json:"achievements"`
}

type GuildRewardItem struct {
	ID                   int                           `json:"id"`
	Name                 string                        `json:"name"`
	Icon                 string                        `json:"icon"`
	Quality              int                           `json:"quality"`
	ItemLevel            int                           `json:"itemLevel"`
	TooltipParams        GuildRewardsItemTooltipParams `json:"tooltipParams"`
	Stats                []GuildRewardsItemStat        `json:"stats"`
	Armor                int                           `json:"armor"`
	Context              string                        `json:"context"`
	BonusLists           []interface{}                 `json:"bonusLists"`
	ArtifactID           int                           `json:"artifactId"`
	DisplayInfoID        int                           `json:"displayInfoId"`
	ArtifactAppearanceID int                           `json:"artifactAppearanceId"`
	ArtifactTraits       []interface{}                 `json:"artifactTraits"`
	Relics               []interface{}                 `json:"relics"`
	Appearance           struct {
	} `json:"appearance"`
}

type GuildRewardsItemTooltipParams struct {
	TimewalkerLevel   int `json:"timewalkerLevel"`
	AzeritePower0     int `json:"azeritePower0"`
	AzeritePower1     int `json:"azeritePower1"`
	AzeritePower2     int `json:"azeritePower2"`
	AzeritePower3     int `json:"azeritePower3"`
	AzeritePowerLevel int `json:"azeritePowerLevel"`
}

type GuildRewardsItemStat struct {
	Stat   int `json:"stat"`
	Amount int `json:"amount"`
}

type GuildRewardsItem struct {
}

type GuildRewardAchievement struct {
	ID          int                   `json:"id"`
	Title       string                `json:"title"`
	Points      int                   `json:"points"`
	Description string                `json:"description"`
	Reward      string                `json:"reward"`
	RewardItems []GuildRewardsItem    `json:"rewardItems"`
	Icon        string                `json:"icon"`
	Criteria    []AchievementCriteria `json:"criteria"`
	AccountWide bool                  `json:"accountWide"`
	FactionID   int                   `json:"factionId"`
}

type GuildRewards struct {
	MinGuildLevel    int                    `json:"minGuildLevel"`
	MinGuildRepLevel int                    `json:"minGuildRepLevel"`
	Achievement      GuildRewardAchievement `json:"achievement,omitempty"`
	Item             GuildRewardItem        `json:"item"`
	Races            []int                  `json:"races,omitempty"`
}

type GuildRewardsData struct {
	Rewards []GuildRewards `json:"rewards"`
}

type GuildPerkSpell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Subtext     string `json:"subtext"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
}

type GuildPerk struct {
	GuildLevel int            `json:"guildLevel"`
	Spell      GuildPerkSpell `json:"spell"`
}

type GuildPerksData struct {
	Perks []GuildPerk `json:"perks"`
}

type GuildAchievementsList struct {
	ID          int                   `json:"id"`
	Title       string                `json:"title"`
	Points      int                   `json:"points"`
	Description string                `json:"description"`
	RewardItems []interface{}         `json:"rewardItems"`
	Icon        string                `json:"icon"`
	Criteria    []AchievementCriteria `json:"criteria"`
	AccountWide bool                  `json:"accountWide"`
	FactionID   int                   `json:"factionId"`
	Reward      string                `json:"reward,omitempty"`
}

type GuildAchievementsCategories struct {
	ID           int                     `json:"id"`
	Achievements []GuildAchievementsList `json:"achievements"`
	Name         string                  `json:"name"`
}

type GuildAchievements struct {
	ID           int                           `json:"id"`
	Achievements []GuildAchievementsList       `json:"achievements"`
	Name         string                        `json:"name"`
	Categories   []GuildAchievementsCategories `json:"categories,omitempty"`
}

type GuildAchievementsData struct {
	Achievements []GuildAchievements `json:"achievements"`
}

type ItemSubclasses struct {
	Subclass int    `json:"subclass"`
	Name     string `json:"name"`
}

type ItemClasses struct {
	Class      int              `json:"class"`
	Name       string           `json:"name"`
	Subclasses []ItemSubclasses `json:"subclasses"`
}

type ItemClassesData struct {
	Classes []ItemClasses `json:"classes"`
}

type TalentNumSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

type TalentNumPetSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

type TalentNumSpell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
}

type TalentNumTalent struct {
	Tier   int              `json:"tier"`
	Column int              `json:"column"`
	Spell  TalentNumSpell   `json:"spell"`
	Spec   TalentNumPetSpec `json:"spec"`
}

type TalentNum struct {
	Specs    []TalentNumSpec       `json:"specs"`
	PetSpecs []TalentNumPetSpec    `json:"petSpecs"`
	Talents  [][][]TalentNumTalent `json:"talents"`
	Class    string                `json:"class"`
}

type TalentsData struct {
	Num1  TalentNum `json:"1"`
	Num2  TalentNum `json:"2"`
	Num3  TalentNum `json:"3"`
	Num4  TalentNum `json:"4"`
	Num5  TalentNum `json:"5"`
	Num6  TalentNum `json:"6"`
	Num7  TalentNum `json:"7"`
	Num8  TalentNum `json:"8"`
	Num9  TalentNum `json:"9"`
	Num10 TalentNum `json:"10"`
	Num11 TalentNum `json:"11"`
	Num12 TalentNum `json:"12"`
}

type PetTypes struct {
	ID              int    `json:"id"`
	Key             string `json:"key"`
	Name            string `json:"name"`
	TypeAbilityID   int    `json:"typeAbilityId"`
	StrongAgainstID int    `json:"strongAgainstId"`
	WeakAgainstID   int    `json:"weakAgainstId"`
}

type PetTypesData struct {
	PetTypes []PetTypes `json:"petTypes"`
}

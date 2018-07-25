package wow

// Battlegroup represents a World of Warcraft battlegroup.
type Battlegroup struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// BattleGroupsData represents World of Warcraft battlegroup data.
type BattleGroupsData struct {
	Battlegroups []Battlegroup `json:"battlegroups"`
}

// CharacterRace represents a World of Warcraft character race.
type CharacterRace struct {
	ID   int    `json:"id"`
	Mask int    `json:"mask"`
	Side string `json:"side"`
	Name string `json:"name"`
}

// CharacterRacesData represents World of Warcraft character races data.
type CharacterRacesData struct {
	Races []CharacterRace `json:"races"`
}

// CharacterClass represents a World of Warcraft character class.
type CharacterClass struct {
	ID        int    `json:"id"`
	Mask      int    `json:"mask"`
	PowerType string `json:"powerType"`
	Name      string `json:"name"`
}

// CharacterClassesData represents World of Warcraft character classes data.
type CharacterClassesData struct {
	Classes []CharacterClass `json:"classes"`
}

// CharacterAchievementCategory represents a World of Warcraft character achievement category.
type CharacterAchievementCategory struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Achievements []Achievement `json:"achievements"`
}

// CharacterAchievementData represents World of Warcraft character achievement data.
type CharacterAchievementData struct {
	ID           int                            `json:"id"`
	Name         string                         `json:"name"`
	Achievements []Achievement                  `json:"achievements,omitempty"`
	Categories   []CharacterAchievementCategory `json:"categories,omitempty"`
}

// CharacterAchievementsData represents World of Warcraft character achievements data.
type CharacterAchievementsData struct {
	Achievements []CharacterAchievementData `json:"achievements"`
}

// GuildRewardItem represents a World of Warcraft guild reward item.
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

// GuildRewardsItemTooltipParams represents tooltip parameters for a World of Warcraft guild reward item.
type GuildRewardsItemTooltipParams struct {
	TimewalkerLevel   int `json:"timewalkerLevel"`
	AzeritePower0     int `json:"azeritePower0"`
	AzeritePower1     int `json:"azeritePower1"`
	AzeritePower2     int `json:"azeritePower2"`
	AzeritePower3     int `json:"azeritePower3"`
	AzeritePowerLevel int `json:"azeritePowerLevel"`
}

// GuildRewardsItemStat represents a World of Warcraft guild reward item stat.
type GuildRewardsItemStat struct {
	Stat   int `json:"stat"`
	Amount int `json:"amount"`
}

// GuildRewardsItem represents a World of Warcraft guild reward item.
type GuildRewardsItem struct {
}

// GuildRewards represents World of Warcraft guild rewards.
type GuildRewards struct {
	MinGuildLevel    int             `json:"minGuildLevel"`
	MinGuildRepLevel int             `json:"minGuildRepLevel"`
	Achievement      Achievement     `json:"achievement,omitempty"`
	Item             GuildRewardItem `json:"item"`
	Races            []int           `json:"races,omitempty"`
}

// GuildRewardsData represents World of Warcraft guild rewards data.
type GuildRewardsData struct {
	Rewards []GuildRewards `json:"rewards"`
}

// GuildPerkSpell represents a World of Warcraft guild perk spell.
type GuildPerkSpell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Subtext     string `json:"subtext"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
}

// GuildPerk represents a World of Warcraft guild perk.
type GuildPerk struct {
	GuildLevel int            `json:"guildLevel"`
	Spell      GuildPerkSpell `json:"spell"`
}

// GuildPerksData represents World of Warcraft guild perks data.
type GuildPerksData struct {
	Perks []GuildPerk `json:"perks"`
}

// GuildAchievementsCategories represents World of Warcraft guild achievements categories.
type GuildAchievementsCategories struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Achievements []Achievement `json:"achievements"`
}

// GuildAchievements represents World of Warcraft guild achievements.
type GuildAchievements struct {
	ID           int                           `json:"id"`
	Name         string                        `json:"name"`
	Achievements []Achievement                 `json:"achievements"`
	Categories   []GuildAchievementsCategories `json:"categories,omitempty"`
}

// GuildAchievementsData represents World of Warcraft guild achievements data.
type GuildAchievementsData struct {
	Achievements []GuildAchievements `json:"achievements"`
}

// ItemSubclasses represents World of Warcraft item subclasses.
type ItemSubclasses struct {
	Subclass int    `json:"subclass"`
	Name     string `json:"name"`
}

// ItemClasses represents World of Warcraft item classes.
type ItemClasses struct {
	Class      int              `json:"class"`
	Name       string           `json:"name"`
	Subclasses []ItemSubclasses `json:"subclasses"`
}

// ItemClassesData represents World of Warcraft item classes data.
type ItemClassesData struct {
	Classes []ItemClasses `json:"classes"`
}

// TalentNumSpec represents a World of Warcraft talent number spec.
type TalentNumSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

// TalentNumPetSpec represents a World of Warcraft talent number pet spec.
type TalentNumPetSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

// TalentNumSpell represents a World of Warcraft talent number spell.
type TalentNumSpell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
}

// TalentNumTalent represents a World of Warcraft talent number talent.
type TalentNumTalent struct {
	Tier   int              `json:"tier"`
	Column int              `json:"column"`
	Spell  TalentNumSpell   `json:"spell"`
	Spec   TalentNumPetSpec `json:"spec"`
}

// TalentNum represents a World of Warcraft talent number.
type TalentNum struct {
	Specs    []TalentNumSpec       `json:"specs"`
	PetSpecs []TalentNumPetSpec    `json:"petSpecs"`
	Talents  [][][]TalentNumTalent `json:"talents"`
	Class    string                `json:"class"`
}

// TalentsData represents World of Warcraft talents data.
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

// PetTypes represents World of Warcraft pet types.
type PetTypes struct {
	ID              int    `json:"id"`
	Key             string `json:"key"`
	Name            string `json:"name"`
	TypeAbilityID   int    `json:"typeAbilityId"`
	StrongAgainstID int    `json:"strongAgainstId"`
	WeakAgainstID   int    `json:"weakAgainstId"`
}

// PetTypesData represents World of Warcraft pet types data.
type PetTypesData struct {
	PetTypes []PetTypes `json:"petTypes"`
}

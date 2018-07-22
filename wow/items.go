package wow

// ItemWeaponDamage represents a World of Warcraft weapon's damage.
type ItemWeaponDamage struct {
	Min      int `json:"min"`
	Max      int `json:"max"`
	ExactMin int `json:"exactMin"`
	ExactMax int `json:"exactMax"`
}

// ItemWeaponInfo represents a World of Warcraft weapon's info.
type ItemWeaponInfo struct {
	Damage      ItemWeaponDamage `json:"damage"`
	WeaponSpeed float64          `json:"weaponSpeed"`
	Dps         float64          `json:"dps"`
}

// BonusStat represents a World of Warcraft bonus statistic.
type BonusStat struct {
	Stat   int `json:"stat"`
	Amount int `json:"amount"`
}

// ItemSource represents a World of Warcraft item source.
type ItemSource struct {
	SourceID   int    `json:"sourceId"`
	SourceType string `json:"sourceType"`
}

// ItemBonusSummary represents a World of Warcraft item bonus summary.
type ItemBonusSummary struct {
	DefaultBonusLists []interface{} `json:"defaultBonusLists"`
	ChanceBonusLists  []interface{} `json:"chanceBonusLists"`
	BonusChances      []interface{} `json:"bonusChances"`
}

// Item represents a World of Warcraft item.
type Item struct {
	ID                     int              `json:"id"`
	DisenchantingSkillRank int              `json:"disenchantingSkillRank"`
	Description            string           `json:"description"`
	Name                   string           `json:"name"`
	Icon                   string           `json:"icon"`
	Stackable              int              `json:"stackable"`
	ItemBind               int              `json:"itemBind"`
	BonusStats             []BonusStat      `json:"bonusStats"`
	ItemSpells             []interface{}    `json:"itemSpells"`
	BuyPrice               int              `json:"buyPrice"`
	ItemClass              int              `json:"itemClass"`
	ItemSubClass           int              `json:"itemSubClass"`
	ContainerSlots         int              `json:"containerSlots"`
	WeaponInfo             ItemWeaponInfo   `json:"weaponInfo"`
	InventoryType          int              `json:"inventoryType"`
	Equippable             bool             `json:"equippable"`
	ItemLevel              int              `json:"itemLevel"`
	MaxCount               int              `json:"maxCount"`
	MaxDurability          int              `json:"maxDurability"`
	MinFactionID           int              `json:"minFactionId"`
	MinReputation          int              `json:"minReputation"`
	Quality                int              `json:"quality"`
	SellPrice              int              `json:"sellPrice"`
	RequiredSkill          int              `json:"requiredSkill"`
	RequiredLevel          int              `json:"requiredLevel"`
	RequiredSkillRank      int              `json:"requiredSkillRank"`
	ItemSource             ItemSource       `json:"itemSource"`
	BaseArmor              int              `json:"baseArmor"`
	HasSockets             bool             `json:"hasSockets"`
	IsAuctionable          bool             `json:"isAuctionable"`
	Armor                  int              `json:"armor"`
	DisplayInfoID          int              `json:"displayInfoId"`
	NameDescription        string           `json:"nameDescription"`
	NameDescriptionColor   string           `json:"nameDescriptionColor"`
	Upgradable             bool             `json:"upgradable"`
	HeroicTooltip          bool             `json:"heroicTooltip"`
	Context                string           `json:"context"`
	BonusLists             []interface{}    `json:"bonusLists"`
	AvailableContexts      []string         `json:"availableContexts"`
	BonusSummary           ItemBonusSummary `json:"bonusSummary"`
	ArtifactID             int              `json:"artifactId"`
}

// ItemSetBonus represents a World of Warcraft item set bonus.
type ItemSetBonus struct {
	Description string `json:"description"`
	Threshold   int    `json:"threshold"`
}

// ItemSet represents a World of Warcraft item set.
type ItemSet struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	SetBonuses []ItemSetBonus `json:"setBonuses"`
	Items      []int          `json:"items"`
}

package d3

// FollowerType represents a Diablo III follower type.
type FollowerType struct {
	TwoHanded bool   `json:"twoHanded"`
	ID        string `json:"id"`
}

// FollowerAttributes represents Diablo III follower attributes.
type FollowerAttributes struct {
	Primary []string `json:"primary"`
}

// FollowerHand represents Diablo III follower hand data.
type FollowerHand struct {
	ID                     string             `json:"id"`
	Name                   string             `json:"name"`
	Icon                   string             `json:"icon"`
	DisplayColor           string             `json:"displayColor"`
	TooltipParams          string             `json:"tooltipParams"`
	RequiredLevel          int                `json:"requiredLevel"`
	ItemLevel              int                `json:"itemLevel"`
	StackSizeMax           int                `json:"stackSizeMax"`
	AccountBound           bool               `json:"accountBound"`
	TypeName               string             `json:"typeName"`
	Type                   FollowerType       `json:"type"`
	Armor                  int                `json:"armor"`
	Damage                 string             `json:"damage"`
	Dps                    string             `json:"dps"`
	AttacksPerSecond       float64            `json:"attacksPerSecond"`
	MinDamage              int                `json:"minDamage"`
	MaxDamage              int                `json:"maxDamage"`
	Slots                  string             `json:"slots"`
	Attributes             FollowerAttributes `json:"attributes"`
	AttributesHTML         FollowerAttributes
	OpenSockets            int    `json:"openSockets"`
	SeasonRequiredToDrop   int    `json:"seasonRequiredToDrop"`
	BlockChance            string `json:"blockChance"`
	IsSeasonRequiredToDrop bool   `json:"isSeasonRequiredToDrop"`
}

// DualWieldFollower represents Diablo III follower dual wielding data.
type DualWieldFollower struct {
	MainHand FollowerHand `json:"mainHand"`
	OffHand  FollowerHand `json:"offHand"`
}

// SingleWieldFollower represents Diablo III follower single wielding data.
type SingleWieldFollower struct {
	MainHand FollowerHand `json:"mainHand"`
}

// DetailedFollowerItems represents detailed Diablo III follower items.
type DetailedFollowerItems struct {
	Templar     DualWieldFollower   `json:"templar"`
	Scoundrel   SingleWieldFollower `json:"scoundrel"`
	Enchantress SingleWieldFollower `json:"enchantress"`
}

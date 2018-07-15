package d3

// Get API Detailed Follower Items
type FollowerType struct {
	TwoHanded bool   `json:"twoHanded"`
	ID        string `json:"id"`
}

type FollowerAttributes struct {
	Primary []string `json:"primary"`
}

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

type DualWieldFollower struct {
	MainHand FollowerHand `json:"mainHand"`
	OffHand  FollowerHand `json:"offHand"`
}

type SingleWieldFollower struct {
	MainHand FollowerHand `json:"mainHand"`
}

type DetailedFollowerItems struct {
	Templar     DualWieldFollower   `json:"templar"`
	Scoundrel   SingleWieldFollower `json:"scoundrel"`
	Enchantress SingleWieldFollower `json:"enchantress"`
}

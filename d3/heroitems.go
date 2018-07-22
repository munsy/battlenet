package d3

// DetailedHeroItemType represents a detailed Diablo III hero item type.
type DetailedHeroItemType struct {
	TwoHanded bool   `json:"twoHanded"`
	ID        string `json:"id"`
}

// DetailedHeroItemAttributes represents a detailed Diablo III hero item attribute.
type DetailedHeroItemAttributes struct {
	Primary   []string `json:"primary"`
	Secondary []string `json:"secondary"`
}

// DetailedHeroItem represents a detailed Diablo III hero item.
type DetailedHeroItem struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Path string `json:"path"`
}

// Gem represents a Diablo III gem.
type Gem struct {
	Item       DetailedHeroItem `json:"item"`
	Attributes []string         `json:"attributes"`
	IsGem      bool             `json:"isGem"`
	IsJewel    bool             `json:"isJewel"`
}

// Dye represents a Diablo III dye.
type Dye struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	TooltipParams string `json:"tooltipParams"`
}

// DetailedHeroItemSlot represents a detailed Diablo III hero item slot.
type DetailedHeroItemSlot struct {
	ID                     string                     `json:"id"`
	Name                   string                     `json:"name"`
	Icon                   string                     `json:"icon"`
	DisplayColor           string                     `json:"displayColor"`
	TooltipParams          string                     `json:"tooltipParams"`
	RequiredLevel          int                        `json:"requiredLevel"`
	ItemLevel              int                        `json:"itemLevel"`
	StackSizeMax           int                        `json:"stackSizeMax"`
	AccountBound           bool                       `json:"accountBound"`
	FlavorText             string                     `json:"flavorText"`
	TypeName               string                     `json:"typeName"`
	Type                   DetailedHeroItemType       `json:"type"`
	Armor                  int                        `json:"armor"`
	AttacksPerSecond       int                        `json:"attacksPerSecond"`
	MinDamage              int                        `json:"minDamage"`
	MaxDamage              int                        `json:"maxDamage"`
	Slots                  string                     `json:"slots"`
	Attributes             DetailedHeroItemAttributes `json:"attributes"`
	AttributesHTML         DetailedHeroItemAttributes `json:"attributesHtml"`
	OpenSockets            int                        `json:"openSockets"`
	Gems                   []Gem                      `json:"gems"`
	SeasonRequiredToDrop   int                        `json:"seasonRequiredToDrop"`
	Dye                    Dye                        `json:"dye"`
	Transmog               HeroItemData               `json:"transmog"`
	IsSeasonRequiredToDrop bool                       `json:"isSeasonRequiredToDrop"`
}

// DetailedHeroItems represents detailed Diablo III hero items.
type DetailedHeroItems struct {
	Head        DetailedHeroItemSlot `json:"head"`
	Neck        DetailedHeroItemSlot `json:"neck"`
	Torso       DetailedHeroItemSlot `json:"torso"`
	Shoulders   DetailedHeroItemSlot `json:"shoulders"`
	Legs        DetailedHeroItemSlot `json:"legs"`
	Waist       DetailedHeroItemSlot `json:"waist"`
	Hands       DetailedHeroItemSlot `json:"hands"`
	Bracers     DetailedHeroItemSlot `json:"bracers"`
	Feet        DetailedHeroItemSlot `json:"feet"`
	LeftFinger  DetailedHeroItemSlot `json:"leftFinger"`
	RightFinger DetailedHeroItemSlot `json:"rightFinger"`
	MainHand    DetailedHeroItemSlot `json:"mainHand"`
	OffHand     DetailedHeroItemSlot `json:"offHand"`
}

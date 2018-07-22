package d3

// ItemMetadata represents Diablo III item metadata.
type ItemMetadata struct {
	TwoHanded bool   `json:"twoHanded"`
	ID        string `json:"id"`
}

// ItemAttributeData represents Diablo III item attribute data.
type ItemAttributeData struct {
	TextHTML string `json:"textHtml"`
	Text     string `json:"text"`
}

// RandomAffix represents a random Diablo III affix.
type RandomAffix struct {
	OneOf []ItemAttributeData `json:"oneOf"`
}

// ItemAttribute represents a Diablo III item attribute.
type ItemAttribute struct {
	Primary   []ItemAttributeData `json:"primary"`
	Secondary []ItemAttributeData `json:"secondary"`
	Other     []interface{}       `json:"other"`
}

// Item represents a Diablo III item.
type Item struct {
	ID                     string        `json:"id"`
	Slug                   string        `json:"slug"`
	Name                   string        `json:"name"`
	Icon                   string        `json:"icon"`
	TooltipParams          string        `json:"tooltipParams"`
	RequiredLevel          int           `json:"requiredLevel"`
	StackSizeMax           int           `json:"stackSizeMax"`
	AccountBound           bool          `json:"accountBound"`
	FlavorText             string        `json:"flavorText"`
	FlavorTextHTML         string        `json:"flavorTextHtml"`
	TypeName               string        `json:"typeName"`
	Type                   ItemMetadata  `json:"type"`
	Damage                 string        `json:"damage"`
	Dps                    string        `json:"dps"`
	DamageHTML             string        `json:"damageHtml"`
	Color                  string        `json:"color"`
	IsSeasonRequiredToDrop bool          `json:"isSeasonRequiredToDrop"`
	SeasonRequiredToDrop   int           `json:"seasonRequiredToDrop"`
	Slots                  []string      `json:"slots"`
	Attributes             ItemAttribute `json:"attributes"`
	RandomAffixes          []RandomAffix `json:"randomAffixes"`
	SetItems               []interface{} `json:"setItems"`
}

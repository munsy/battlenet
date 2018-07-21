package d3

type ItemMetadata struct {
	TwoHanded bool   `json:"twoHanded"`
	ID        string `json:"id"`
}

type ItemAttributeData struct {
	TextHTML string `json:"textHtml"`
	Text     string `json:"text"`
}

type RandomAffix struct {
	OneOf []ItemAttributeData `json:"oneOf"`
}

type ItemAttribute struct {
	Primary   []ItemAttributeData `json:"primary"`
	Secondary []ItemAttributeData `json:"secondary"`
	Other     []interface{}       `json:"other"`
}

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

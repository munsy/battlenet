package wow

type Realm struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type BonusList struct {
	BonusListID int `json:"bonusListId"`
}

type Modifier struct {
	Type  int `json:"type"`
	Value int `json:"value"`
}

type Auction struct {
	Auc          int         `json:"auc"`
	Item         int         `json:"item"`
	Owner        string      `json:"owner"`
	OwnerRealm   string      `json:"ownerRealm"`
	Bid          int         `json:"bid"`
	Buyout       int         `json:"buyout"`
	Quantity     int         `json:"quantity"`
	TimeLeft     string      `json:"timeLeft"`
	Rand         int         `json:"rand"`
	Seed         int         `json:"seed"`
	Context      int         `json:"context"`
	BonusLists   []BonusList `json:"bonusLists,omitempty"`
	Modifiers    []Modifier  `json:"modifiers,omitempty"`
	PetSpeciesID int         `json:"petSpeciesId,omitempty"`
	PetBreedID   int         `json:"petBreedId,omitempty"`
	PetLevel     int         `json:"petLevel,omitempty"`
	PetQualityID int         `json:"petQualityId,omitempty"`
}

type File struct {
	URL          string `json:"url"`
	LastModified int64  `json:"lastModified"`
}

type AuctionFileQueryData struct {
	Files []File `json:"files"`
}

type AuctionJSONFileData struct {
	Realms   []Realm   `json:"realms"`
	Auctions []Auction `json:"auctions"`
}

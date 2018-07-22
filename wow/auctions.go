package wow

// AuctionRealm represents the World of Warcraft realm that the given auction is taking place on.
type AuctionRealm struct {
	Name            string   `json:"name"`
	Slug            string   `json:"slug"`
	Battlegroup     string   `json:"battlegroup"`
	Locale          string   `json:"locale"`
	Timezone        string   `json:"timezone"`
	ConnectedRealms []string `json:"connected_realms"`
}

// BonusList represents a World of Warcraft auction bonus list.
type BonusList struct {
	BonusListID int `json:"bonusListId"`
}

// Modifier holds data about a World of Warcraft auction modifier.
type Modifier struct {
	Type  int `json:"type"`
	Value int `json:"value"`
}

// Auction contains data about a specific World of Warcraft auction.
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

// File contains World of Warcraft auction data for further parsing.
type File struct {
	URL          string `json:"url"`
	LastModified int64  `json:"lastModified"`
}

// AuctionFileQueryData contains a list of World of Warcraft auction data files.
type AuctionFileQueryData struct {
	Files []File `json:"files"`
}

// AuctionJSONFileData is the JSON representation of a World of Warcraft auction data file's contents.
type AuctionJSONFileData struct {
	Realms   []AuctionRealm `json:"realms"`
	Auctions []Auction      `json:"auctions"`
}

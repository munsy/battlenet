package wow

// Realm represents a World of Warcraft realm.
type Realm struct {
	Type            string   `json:"type"`
	Population      string   `json:"population"`
	Queue           bool     `json:"queue"`
	Status          bool     `json:"status"`
	Name            string   `json:"name"`
	Slug            string   `json:"slug"`
	Battlegroup     string   `json:"battlegroup"`
	Locale          string   `json:"locale"`
	Timezone        string   `json:"timezone"`
	ConnectedRealms []string `json:"connected_realms"`
}

// RealmStatus represents the status of every World of Warcraft realm.
type RealmStatus struct {
	Realms []Realm `json:"realms"`
}

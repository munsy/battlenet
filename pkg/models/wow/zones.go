package wow

// Location represents the location of a World of Warcraft object.
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ZoneBossNPC represents an NPC accompanying a World of Warcraft boss in a zone.
type ZoneBossNPC struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	URLSlug           string `json:"urlSlug"`
	CreatureDisplayID int    `json:"creatureDisplayId"`
}

// ZoneBoss represents a World of Warcraft boss in a zone.
type ZoneBoss struct {
	ID                    int           `json:"id"`
	Name                  string        `json:"name"`
	URLSlug               string        `json:"urlSlug"`
	Description           string        `json:"description"`
	ZoneID                int           `json:"zoneId"`
	AvailableInNormalMode bool          `json:"availableInNormalMode"`
	AvailableInHeroicMode bool          `json:"availableInHeroicMode"`
	Health                int           `json:"health"`
	HeroicHealth          int           `json:"heroicHealth"`
	Level                 int           `json:"level"`
	HeroicLevel           int           `json:"heroicLevel"`
	JournalID             int           `json:"journalId"`
	Npcs                  []ZoneBossNPC `json:"npcs"`
}

// Zone contains data about a specific World of Warcraft zone.
type Zone struct {
	ID                    int        `json:"id"`
	Name                  string     `json:"name"`
	URLSlug               string     `json:"urlSlug"`
	Description           string     `json:"description"`
	Location              Location   `json:"location"`
	ExpansionID           int        `json:"expansionId"`
	Patch                 string     `json:"patch,omitempty"`
	NumPlayers            string     `json:"numPlayers"`
	IsDungeon             bool       `json:"isDungeon"`
	IsRaid                bool       `json:"isRaid"`
	AdvisedMinLevel       int        `json:"advisedMinLevel"`
	AdvisedMaxLevel       int        `json:"advisedMaxLevel"`
	AdvisedHeroicMinLevel int        `json:"advisedHeroicMinLevel"`
	AdvisedHeroicMaxLevel int        `json:"advisedHeroicMaxLevel"`
	AvailableModes        []string   `json:"availableModes"`
	LfgNormalMinGearLevel int        `json:"lfgNormalMinGearLevel"`
	LfgHeroicMinGearLevel int        `json:"lfgHeroicMinGearLevel"`
	Floors                int        `json:"floors"`
	Bosses                []ZoneBoss `json:"bosses"`
}

// ZoneMasterList contains all of the possible World of Warcraft zones.
type ZoneMasterList struct {
	Zones []Zone `json:"zones"`
}

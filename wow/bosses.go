package wow

// BossNPC represents a World of Warcraft NPC related to a boss.
type BossNPC struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	URLSlug           string `json:"urlSlug"`
	CreatureDisplayID int    `json:"creatureDisplayId"`
}

// BossLocation represents a World of Warcraft boss location.
type BossLocation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Boss represents a World of Warcraft boss.
type Boss struct {
	ID                    int          `json:"id"`
	Name                  string       `json:"name"`
	URLSlug               string       `json:"urlSlug"`
	Description           string       `json:"description,omitempty"`
	ZoneID                int          `json:"zoneId"`
	AvailableInNormalMode bool         `json:"availableInNormalMode"`
	AvailableInHeroicMode bool         `json:"availableInHeroicMode"`
	Health                int          `json:"health"`
	HeroicHealth          int          `json:"heroicHealth"`
	Level                 int          `json:"level"`
	HeroicLevel           int          `json:"heroicLevel"`
	JournalID             int          `json:"journalId,omitempty"`
	Npcs                  []BossNPC    `json:"npcs"`
	Location              BossLocation `json:"location,omitempty"`
}

// BossMasterList represents the World of Warcraft boss master list.
type BossMasterList struct {
	Bosses []Boss `json:"bosses"`
}

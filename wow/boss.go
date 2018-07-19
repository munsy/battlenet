package wow

type BossNPC struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	URLSlug           string `json:"urlSlug"`
	CreatureDisplayID int    `json:"creatureDisplayId"`
}

type BossLocation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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

type BossMasterList struct {
	Bosses []Boss `json:"bosses"`
}

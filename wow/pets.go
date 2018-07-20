package wow

type PetStats struct {
	SpeciesID    int `json:"speciesId"`
	BreedID      int `json:"breedId"`
	PetQualityID int `json:"petQualityId"`
	Level        int `json:"level"`
	Health       int `json:"health"`
	Power        int `json:"power"`
	Speed        int `json:"speed"`
}

type Pet struct {
	CanBattle     bool     `json:"canBattle"`
	CreatureID    int      `json:"creatureId"`
	Name          string   `json:"name"`
	Family        string   `json:"family"`
	Icon          string   `json:"icon"`
	QualityID     int      `json:"qualityId"`
	Stats         PetStats `json:"stats"`
	StrongAgainst []string `json:"strongAgainst"`
	TypeID        int      `json:"typeId"`
	WeakAgainst   []string `json:"weakAgainst"`
}

type PetData struct {
	Pets []Pet `json:"pets"`
}

type PetAbility struct {
	Slot          int    `json:"slot"`
	Order         int    `json:"order"`
	RequiredLevel int    `json:"requiredLevel"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	Cooldown      int    `json:"cooldown"`
	Rounds        int    `json:"rounds"`
	PetTypeID     int    `json:"petTypeId"`
	IsPassive     bool   `json:"isPassive"`
	HideHints     bool   `json:"hideHints"`
}

type PetSpecies struct {
	SpeciesID   int          `json:"speciesId"`
	PetTypeID   int          `json:"petTypeId"`
	CreatureID  int          `json:"creatureId"`
	Name        string       `json:"name"`
	CanBattle   bool         `json:"canBattle"`
	Icon        string       `json:"icon"`
	Description string       `json:"description"`
	Source      string       `json:"source"`
	Abilities   []PetAbility `json:"abilities"`
}

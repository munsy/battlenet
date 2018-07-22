package d3

// Kills represents all of a Diablo III character's kills.
type Kills struct {
	Monsters         int `json:"monsters"`
	Elites           int `json:"elites"`
	HardcoreMonsters int `json:"hardcoreMonsters"`
}

// AccountHero represents a Diablo III hero.
type AccountHero struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	ClassSlug    string `json:"classSlug"`
	Gender       int    `json:"gender"`
	Level        int    `json:"level"`
	Kills        Kills  `json:"kills"`
	ParagonLevel int    `json:"paragonLevel"`
	Hardcore     bool   `json:"hardcore"`
	Seasonal     bool   `json:"seasonal"`
	Dead         bool   `json:"dead"`
	LastUpdated  int    `json:"last-updated"`
}

// TimePlayed represents time played on Diablo III.
type TimePlayed struct {
	DemonHunter float64 `json:"demon-hunter"`
	Barbarian   float64 `json:"barbarian"`
	WitchDoctor float64 `json:"witch-doctor"`
	Necromancer float64 `json:"necromancer"`
	Wizard      float64 `json:"wizard"`
	Monk        int     `json:"monk"`
	Crusader    float64 `json:"crusader"`
}

// Progression represents Diablo III character progression.
type Progression struct {
	Act1 bool `json:"act1"`
	Act3 bool `json:"act3"`
	Act2 bool `json:"act2"`
	Act5 bool `json:"act5"`
	Act4 bool `json:"act4"`
}

// Death represents a Diablo III death.
type Death struct {
	Killer int `json:"killer"`
	Time   int `json:"time"`
}

// FallenHero represents a fallen Diablo III hero.
type FallenHero struct {
	HeroID   int    `json:"heroId"`
	Name     string `json:"name"`
	Class    string `json:"class"`
	Level    int    `json:"level"`
	Elites   int    `json:"elites"`
	Hardcore bool   `json:"hardcore"`
	Death    Death  `json:"death"`
	Gender   int    `json:"gender"`
}

// Season represents a Diablo III season.
type Season struct {
	SeasonID             int        `json:"seasonId"`
	ParagonLevel         int        `json:"paragonLevel"`
	ParagonLevelHardcore int        `json:"paragonLevelHardcore"`
	Kills                Kills      `json:"kills"`
	TimePlayed           TimePlayed `json:"timePlayed"`
	HighestHardcoreLevel int        `json:"highestHardcoreLevel"`
}

// SeasonalProfiles represents a Diablo III seasonal profile.
type SeasonalProfiles struct {
	Season14 Season `json:"season14"`
	Season13 Season `json:"season13"`
	Season12 Season `json:"season12"`
	Season11 Season `json:"season11"`
	Season10 Season `json:"season10"`
	Season9  Season `json:"Season9"`
	Season8  Season `json:"season8"`
	Season7  Season `json:"season7"`
	Season6  Season `json:"season6"`
	Season5  Season `json:"season5"`
	Season4  Season `json:"season4"`
	Season3  Season `json:"season3"`
	Season2  Season `json:"season2"`
	Season1  Season `json:"season1"`
	Season0  Season `json:"season0"`
}

// Profession represents a Diablo III profession.
type Profession struct {
	Slug  string `json:"slug"`
	Level int    `json:"level"`
}

// Account represents a Diablo III account.
type Account struct {
	BattleTag                  string           `json:"battleTag"`
	ParagonLevel               int              `json:"paragonLevel"`
	ParagonLevelHardcore       int              `json:"paragonLevelHardcore"`
	ParagonLevelSeason         int              `json:"paragonLevelSeason"`
	ParagonLevelSeasonHardcore int              `json:"paragonLevelSeasonHardcore"`
	GuildName                  string           `json:"guildName"`
	Heroes                     []AccountHero    `json:"heroes"`
	LastHeroPlayed             int              `json:"lastHeroPlayed"`
	LastUpdated                int              `json:"lastUpdated"`
	Kills                      Kills            `json:"kills"`
	HighestHardcoreLevel       int              `json:"highestHardcoreLevel"`
	TimePlayed                 TimePlayed       `json:"timePlayed"`
	Progression                Progression      `json:"progression"`
	FallenHeroes               []FallenHero     `json:"fallenHeroes"`
	SeasonalProfiles           SeasonalProfiles `json:"seasonalProfiles"`
	Blacksmith                 Profession       `json:"blacksmith"`
	Jeweler                    Profession       `json:"jeweler"`
	Mystic                     Profession       `json:"mystic"`
	BlacksmithSeason           Profession       `json:"blacksmithSeason"`
	JewelerSeason              Profession       `json:"jewelerSeason"`
	MysticSeason               Profession       `json:"mysticSeason"`
	BlacksmithHardcore         Profession       `json:"blacksmithHardcore"`
	JewelerHardcore            Profession       `json:"jewelerHardcore"`
	MysticHardcore             Profession       `json:"mysticHardcore"`
	BlacksmithSeasonHardcore   Profession       `json:"blacksmithSeasonHardcore"`
	JewelerSeasonHardcore      Profession       `json:"jewelerSeasonHardcore"`
	MysticSeasonHardcore       Profession       `json:"mysticSeasonHardcore"`
}

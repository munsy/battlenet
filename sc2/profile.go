package sc2

type Career struct {
	PrimaryRace      string `json:"primaryRace"`
	TerranWins       int    `json:"terranWins"`
	ProtossWins      int    `json:"protossWins"`
	ZergWins         int    `json:"zergWins"`
	SeasonTotalGames int    `json:"seasonTotalGames"`
	CareerTotalGames int    `json:"careerTotalGames"`
}

type Terran struct {
	Level          int `json:"level"`
	TotalLevelXP   int `json:"totalLevelXP"`
	CurrentLevelXP int `json:"currentLevelXP"`
}

type Zerg struct {
	Level          int `json:"level"`
	TotalLevelXP   int `json:"totalLevelXP"`
	CurrentLevelXP int `json:"currentLevelXP"`
}

type Protoss struct {
	Level          int `json:"level"`
	TotalLevelXP   int `json:"totalLevelXP"`
	CurrentLevelXP int `json:"currentLevelXP"`
}

type SwarmLevels struct {
	Level   int     `json:"level"`
	Terran  Terran  `json:"terran"`
	Zerg    Zerg    `json:"zerg"`
	Protoss Protoss `json:"protoss"`
}

type Portrait struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
	Offset int    `json:"offset"`
	URL    string `json:"url"`
}

type Campaign struct {
	Wol string `json:"wol"`
}

type Season struct {
	SeasonID             int `json:"seasonId"`
	SeasonNumber         int `json:"seasonNumber"`
	SeasonYear           int `json:"seasonYear"`
	TotalGamesThisSeason int `json:"totalGamesThisSeason"`
}

type Rewards struct {
	Selected []interface{} `json:"selected"`
	Earned   []int64       `json:"earned"`
}

type CategoryPoints struct {
	Num4325377 int `json:"4325377"`
	Num4325379 int `json:"4325379"`
	Num4325410 int `json:"4325410"`
	Num4330138 int `json:"4330138"`
	Num4364473 int `json:"4364473"`
	Num4386911 int `json:"4386911"`
}

type Points struct {
	TotalPoints    int            `json:"totalPoints"`
	CategoryPoints CategoryPoints `json:"categoryPoints"`
}

type Achievement struct {
	AchievementID  int64 `json:"achievementId"`
	CompletionDate int   `json:"completionDate"`
}

type Achievements struct {
	Points       Points        `json:"points"`
	Achievements []Achievement `json:"achievements"`
}

type Profile struct {
	ID           int          `json:"id"`
	Realm        int          `json:"realm"`
	DisplayName  string       `json:"displayName"`
	ClanName     string       `json:"clanName"`
	ClanTag      string       `json:"clanTag"`
	ProfilePath  string       `json:"profilePath"`
	Portrait     Portrait     `json:"portrait"`
	Career       Career       `json:"career"`
	SwarmLevels  SwarmLevels  `json:"swarmLevels"`
	Campaign     Campaign     `json:"campaign"`
	Season       Season       `json:"season"`
	Rewards      Rewards      `json:"rewards"`
	Achievements Achievements `json:"achievements"`
}

package sc2

// Ladder represents a Starcraft II ladder.
type Career struct {
	PrimaryRace      string `json:"primaryRace"`
	TerranWins       int    `json:"terranWins"`
	ProtossWins      int    `json:"protossWins"`
	ZergWins         int    `json:"zergWins"`
	SeasonTotalGames int    `json:"seasonTotalGames"`
	CareerTotalGames int    `json:"careerTotalGames"`
}

// Ladder represents a Starcraft II ladder.
type Terran struct {
	Level          int `json:"level"`
	TotalLevelXP   int `json:"totalLevelXP"`
	CurrentLevelXP int `json:"currentLevelXP"`
}

// Ladder represents a Starcraft II ladder.
type Zerg struct {
	Level          int `json:"level"`
	TotalLevelXP   int `json:"totalLevelXP"`
	CurrentLevelXP int `json:"currentLevelXP"`
}

// Ladder represents a Starcraft II ladder.
type Protoss struct {
	Level          int `json:"level"`
	TotalLevelXP   int `json:"totalLevelXP"`
	CurrentLevelXP int `json:"currentLevelXP"`
}

// Ladder represents a Starcraft II ladder.
type SwarmLevels struct {
	Level   int     `json:"level"`
	Terran  Terran  `json:"terran"`
	Zerg    Zerg    `json:"zerg"`
	Protoss Protoss `json:"protoss"`
}

// Icon represents a Starcraft II icon.
type Icon struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
	Offset int    `json:"offset"`
	URL    string `json:"url"`
}

// Campaign represents a Starcraft II campaign.
type Campaign struct {
	Wol string `json:"wol"`
}

// Season represents a Starcraft II season.
type Season struct {
	SeasonID             int `json:"seasonId"`
	SeasonNumber         int `json:"seasonNumber"`
	SeasonYear           int `json:"seasonYear"`
	TotalGamesThisSeason int `json:"totalGamesThisSeason"`
}

// Rewards represents Starcraft II rewards.
type Rewards struct {
	Selected []interface{} `json:"selected"`
	Earned   []int64       `json:"earned"`
}

// CategoryPoints represents Starcraft II category points.
type CategoryPoints interface{}

// Points represents Starcraft II points.
type Points struct {
	TotalPoints    int            `json:"totalPoints"`
	CategoryPoints CategoryPoints `json:"categoryPoints"`
}

// Achievement represents a Starcraft II achievement.
type Achievement struct {
	AchievementID  int64 `json:"achievementId"`
	CompletionDate int   `json:"completionDate"`
}

// Achievements represents Starcraft II achievements.
type Achievements struct {
	Points       Points        `json:"points"`
	Achievements []Achievement `json:"achievements"`
}

// Character represents a Starcraft II character.
type Character struct {
	ID           int          `json:"id"`
	Realm        int          `json:"realm"`
	DisplayName  string       `json:"displayName"`
	ClanName     string       `json:"clanName"`
	ClanTag      string       `json:"clanTag"`
	ProfilePath  string       `json:"profilePath"`
	Portrait     Icon         `json:"portrait"`
	Career       Career       `json:"career"`
	SwarmLevels  SwarmLevels  `json:"swarmLevels"`
	Campaign     Campaign     `json:"campaign"`
	Season       Season       `json:"season"`
	Rewards      Rewards      `json:"rewards"`
	Achievements Achievements `json:"achievements"`
}

// LadderData represents Starcraft II ladder data.
type LadderData struct {
	LadderName       string `json:"ladderName"`
	LadderID         int    `json:"ladderId"`
	Division         int    `json:"division"`
	Rank             int    `json:"rank"`
	League           string `json:"league"`
	MatchMakingQueue string `json:"matchMakingQueue"`
	Wins             int    `json:"wins"`
	Losses           int    `json:"losses"`
	Showcase         bool   `json:"showcase"`
}

// SeasonData represents Starcraft II ladder season data.
type SeasonData struct {
	Ladder     []LadderData  `json:"ladder"`
	Characters []Character   `json:"characters"`
	NonRanked  []interface{} `json:"nonRanked"`
}

// LadderSeasons represents Starcraft II ladder seasons.
type LadderSeasons struct {
	CurrentSeason     []SeasonData  `json:"currentSeason"`
	PreviousSeason    []SeasonData  `json:"previousSeason"`
	ShowcasePlacement []interface{} `json:"showcasePlacement"`
}

// Match represents a Starcraft II match.
type Match struct {
	Map      string `json:"map"`
	Type     string `json:"type"`
	Decision string `json:"decision"`
	Speed    string `json:"speed"`
	Date     int    `json:"date"`
}

// MatchHistory represents Starcraft II match history.
type MatchHistory struct {
	Matches []Match `json:"matches"`
}

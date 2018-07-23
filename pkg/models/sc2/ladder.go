package sc2

// LadderMember represents a Starcraft II character's ladder data.
type LadderMember struct {
	Character      Character `json:"character"`
	JoinTimestamp  int       `json:"joinTimestamp"`
	Points         int       `json:"points"`
	Wins           int       `json:"wins"`
	Losses         int       `json:"losses"`
	HighestRank    int       `json:"highestRank"`
	PreviousRank   int       `json:"previousRank"`
	FavoriteRaceP1 string    `json:"favoriteRaceP1"`
}

// Ladder contains all of the data for a Starcraft II ladder.
type Ladder struct {
	LadderMembers []LadderMember `json:"ladderMembers"`
}

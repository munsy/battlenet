package sc2

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

type Ladder struct {
	LadderMembers []LadderMember `json:"ladderMembers"`
}

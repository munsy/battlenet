package wow

// GuildProfile contains all the information for a guild's members,
// Battlegroup, achievement points, emblem, etc.
type GuildProfile struct {
	LastModified      int             `json:"lastModified"`
	Name              string          `json:"name"`
	Realm             string          `json:"realm"`
	Battlegroup       string          `json:"battlegroup"`
	Level             int             `json:"level"`
	Side              int             `json:"side"`
	AchievementPoints int             `json:"achievementPoints"`
	Members           []GuildMember   `json:"members"`
	Emblem            EmblemInfo      `json:"emblem"`
	News              []GuildNews     `json:"news"`
	Challenge         []ChallengeData `json:"challenge"`
}

// GuildNews represents news for a World of Warcraft guild.
type GuildNews struct {
	Type        string      `json:"type"`
	Character   string      `json:"character"`
	Timestamp   int64       `json:"timestamp"`
	ItemID      int         `json:"itemId,omitempty"`
	Context     string      `json:"context"`
	BonusLists  []int       `json:"bonusLists"`
	Achievement Achievement `json:"achievement,omitempty"`
}

// MapRankingCriteria represents ranking criteria for a World of Warcraft map.
type MapRankingCriteria struct {
	Time         int  `json:"time"`
	Hours        int  `json:"hours"`
	Minutes      int  `json:"minutes"`
	Seconds      int  `json:"seconds"`
	Milliseconds int  `json:"milliseconds"`
	IsPositive   bool `json:"isPositive"`
}

// MapData represents data about a World of Warcraft map.
type MapData struct {
	ID               int                `json:"id"`
	Name             string             `json:"name"`
	Slug             string             `json:"slug"`
	HasChallengeMode bool               `json:"hasChallengeMode"`
	BronzeCriteria   MapRankingCriteria `json:"bronzeCriteria"`
	SilverCriteria   MapRankingCriteria `json:"silverCriteria"`
	GoldCriteria     MapRankingCriteria `json:"goldCriteria"`
}

// ChallengeData represents data about a World of Warcraft challenge.
type ChallengeData struct {
	Realm  Realm         `json:"realm"`
	Map    MapData       `json:"map"`
	Groups []interface{} `json:"groups"`
}

// EmblemInfo contains information about a guild's emblem.
type EmblemInfo struct {
	Icon              int    `json:"icon"`
	IconColor         string `json:"iconColor"`
	IconColorID       int    `json:"iconColorId"`
	Border            int    `json:"border"`
	BorderColor       string `json:"borderColor"`
	BorderColorID     int    `json:"borderColorId"`
	BackgroundColor   string `json:"backgroundColor"`
	BackgroundColorID int    `json:"backgroundColorId"`
}

// GuildMember represents a list of characters that are a member
// of the guild. When the members list is requested, a list of character
// objects is returned. Each object in the returned members list
// contains a character block as well as a rank field.
type GuildMember struct {
	Character GuildMemberInfo `json:"character"`
	Rank      int             `json:"rank"`
}

// GuildMemberInfo gives information for a member of the guild.
// Contains a character block as well as a rank field.
type GuildMemberInfo struct {
	Name              string        `json:"name"`
	Realm             string        `json:"realm"`
	Battlegroup       string        `json:"battlegroup"`
	Class             int           `json:"class"`
	Race              int           `json:"race"`
	Gender            int           `json:"gender"`
	Level             int           `json:"level"`
	AchievementPoints int           `json:"achievementPoints"`
	Thumbnail         string        `json:"thumbnail"`
	Spec              CharacterSpec `json:"spec"`
	Guild             string        `json:"guild"`
	GuildRealm        string        `json:"guildRealm"`
	LastModified      int           `json:"lastModified"`
}

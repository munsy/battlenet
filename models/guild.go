package models

// https://us.api.battle.net/wow/guild/stormrage/Casual%20Cupcake?fields=members&locale=en_US&apikey= + os.Getenv("BNETKEY")

// GuildInfo contains all the information for a guild's members,
// Battlegroup, achievement points, emblem, etc.
type GuildInfo struct {
	LastModified		int		`json:"lastModified"`
	Name			string		`json:"name"`
	Realm			string		`json:"realm"`
	Battlegroup		string		`json:"battlegroup"`
	Level			int		`json:"level"`
	Side			int		`json:"side"`
	AchievementPoints	int		`json:"achievementPoints"`
	Members			GuildMembers	`json:"members"`
	Emblem			EmblemInfo	`json:"emblem"`
}

// EmblemInfo contains information about a guild's emblem.
type EmblemInfo struct {
	Icon			int	`json:"icon"`
	IconColor		string	`json:"iconColor"`
	IconColorId		int	`json:"iconColorId"`
	Border			int	`json:"border"`
	BorderColor		string	`json:"borderColor"`
	BorderColorId		int	`json:"borderColorId"`
	BackgroundColor		string	`json:"backgroundColor"`
	BackgroundColorId	int	`json:"backgroundColorId"`
}

// GuildMember contains a character block as well as a rank field.
type GuildMember struct {
	Character	GuildMemberInfo	`json:"character"`
	Rank		int		`json:"rank"`
}

// A list of characters that are a member of the guild. When the 
// members list is requested, a list of character objects is 
// returned. Each object in the returned members list contains 
// a character block as well as a rank field.
type GuildMembers []GuildMember

// GuildMemberInfo gives information for a member of the guild. 
// Contains a character block as well as a rank field.
type GuildMemberInfo struct {
	Name			string		`json:"name"`
	Realm			string		`json:"realm"`
	Battlegroup		string		`json:"battlegroup"`
	Class			int		`json:"class"`
	Race			int		`json:"race"`
	Gender			int		`json:"gender"`
	Level			int		`json:"level"`
	AchievementPoints	int		`json:"achievementPoints"`
	Thumbnail		string		`json:"thumbnail"`
	Spec			CharacterSpec	`json:"spec"`
	Guild			string		`json:"guild"`
	GuildRealm		string		`json:"guildRealm"`
	LastModified		int		`json:"lastModified"`
}


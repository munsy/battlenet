package models

// CharacterSpec represents the character specilization information 
// for a particular WoW character from a given WoW profile.
type CharacterSpec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

// Character represents the character information from a particular WoW profile.
type Character struct {
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

// Characters type represents a Character slice.
type Characters struct{
	CharacterList []Character `json:"characters"`
}


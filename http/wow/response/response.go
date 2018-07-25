package response

import (
	"github.com/munsy/gobattlenet/models/wow"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

// Achievement contains World of Warcraft Achievement response data.
type Achievement struct {
	Data     *wow.Achievement
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// AuctionJSONFile contains World of Warcraft AuctionJSONFile response data.
type AuctionJSONFile struct {
	Data     *wow.AuctionJSONFileData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// BossMasterList contains World of Warcraft BossMasterList response data.
type BossMasterList struct {
	Data     *wow.BossMasterList
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Boss contains World of Warcraft Boss response data.
type Boss struct {
	Data     *wow.Boss
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Character contains World of Warcraft Character response data.
type Character struct {
	Data     *wow.CharacterData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// GuildProfile contains World of Warcraft GuildProfile response data.
type GuildProfile struct {
	Data     *wow.GuildProfile
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Item contains World of Warcraft Item response data.
type Item struct {
	Data     *wow.Item
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// ItemSet contains World of Warcraft ItemSet response data.
type ItemSet struct {
	Data     *wow.ItemSet
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// MountList contains World of Warcraft MountList response data.
type MountList struct {
	Data     *wow.MountList
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Pet contains World of Warcraft Pet response data.
type Pet struct {
	Data     *wow.PetData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// PetAbility contains World of Warcraft PetAbility response data.
type PetAbility struct {
	Data     *wow.PetAbility
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// PetSpecies contains World of Warcraft PetSpecies response data.
type PetSpecies struct {
	Data     *wow.PetSpecies
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// PetStats contains World of Warcraft PetStats response data.
type PetStats struct {
	Data     *wow.PetStats
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Quest contains World of Warcraft Quest response data.
type Quest struct {
	Data     *wow.Quest
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// RealmStatus contains World of Warcraft RealmStatus response data.
type RealmStatus struct {
	Data     *wow.RealmStatus
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Recipe contains World of Warcraft Recipe response data.
type Recipe struct {
	Data     *wow.Recipe
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Spell contains World of Warcraft Spell response data.
type Spell struct {
	Data     *wow.Spell
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// ZoneMasterList contains World of Warcraft ZoneMasterList response data.
type ZoneMasterList struct {
	Data     *wow.ZoneMasterList
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Zone contains World of Warcraft Zone response data.
type Zone struct {
	Data     *wow.Zone
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// BattleGroups contains World of Warcraft BattleGroups response data.
type BattleGroups struct {
	Data     *wow.BattleGroupsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// CharacterRaces contains World of Warcraft CharacterRaces response data.
type CharacterRaces struct {
	Data     *wow.CharacterRacesData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// CharacterClasses contains World of Warcraft CharacterClasses response data.
type CharacterClasses struct {
	Data     *wow.CharacterClassesData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// CharacterAchievements contains World of Warcraft CharacterAchievements response data.
type CharacterAchievements struct {
	Data     *wow.CharacterAchievementsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// GuildRewards contains World of Warcraft GuildRewards response data.
type GuildRewards struct {
	Data     *wow.GuildRewardsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// GuildPerks contains World of Warcraft GuildPerks response data.
type GuildPerks struct {
	Data     *wow.GuildPerksData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// GuildAchievements contains World of Warcraft GuildAchievements response data.
type GuildAchievements struct {
	Data     *wow.GuildAchievementsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// ItemClasses contains World of Warcraft ItemClasses response data.
type ItemClasses struct {
	Data     *wow.ItemClassesData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Talents contains World of Warcraft Talents response data.
type Talents struct {
	Data     *wow.TalentsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// PetTypes contains World of Warcraft PetTypes response data.
type PetTypes struct {
	Data     *wow.PetTypesData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

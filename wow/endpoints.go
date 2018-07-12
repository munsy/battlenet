package wow

import (
	"strconv"

	"github.com/munsy/gobattlenet/regions"
)

var (
	wowAPI = regions.API() + "wow/"

	//	PROFILE
	User         = wowAPI + "user/"
	OauthProfile = User + "characters" // WOW OAUTH PROFILE /WOW/USER/CHARACTERS

	//	ACHIEVEMENT
	Achievement = func(id int) string { return wowAPI + "achievement/" + strconv.Itoa(id) } // ACHIEVEMENT /WOW/ACHIEVEMENT/:ID

	//	AUCTION
	endpointAuction   = wowAPI + "auction/"
	AuctionDataStatus = func(realm string) string { return endpointAuction + "data/" + realm } // AUCTION DATA STATUS /WOW/AUCTION/DATA/:REALM

	//	BOSS
	BossMasterList = wowAPI + "boss/"                                                 // MASTER LIST /WOW/BOSS/
	BossInfo       = func(id int) string { return BossMasterList + strconv.Itoa(id) } // BOSS /WOW/BOSS/:BOSSID

	//	CHALLENGE MODE
	endpointChallenge = wowAPI + "challenge/"
	RealmLeaderboard  = func(realm string) string { return endpointChallenge + realm } // REALM LEADERBOARD /WOW/CHALLENGE/:REALM
	RegionLeaderboard = endpointChallenge + "region"                                   // REGION LEADERBOARD /WOW/CHALLENGE/REGION

	//	CHARACTER PROFILE
	endpointCharacter = wowAPI + "character/"
	CharacterProfile  = func(realm, characterName string) string { return endpointCharacter + realm + "/" + characterName } // CHARACTER PROFILE 	/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// ACHIEVEMENTS 	/WOW/CHARACTER/:REALM/:CHARACTERNAME // Implement the rest differently?
	// APPEARANCE 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// FEED 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// GUILD 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// HUNTER PETS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// ITEMS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// MOUNTS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// PETS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// PET SLOTS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// PROFESSIONS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// PROGRESSION 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// PVP 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// QUESTS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// REPUTATION 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// STATISTICS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// STATS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// TALENTS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// TITLES 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	// AUDIT 			/WOW/CHARACTER/:REALM/:CHARACTERNAME

	//	GUILD PROFILE
	endpointGuild = wowAPI + "guild/"
	GuildProfile  = func(realm, guildName string) string { return endpointGuild + realm + "/" + guildName } // GUILD PROFILE 	/WOW/GUILD/:REALM/:GUILDNAME
	// MEMBERS 			/WOW/GUILD/:REALM/:GUILDNAME // implement the rest differently?
	// ACHIEVEMENTS 	/WOW/GUILD/:REALM/:GUILDNAME
	// NEWS 			/WOW/GUILD/:REALM/:GUILDNAME
	// CHALLENGE 		/WOW/GUILD/:REALM/:GUILDNAME

	//	ITEM
	endpointItem = wowAPI + "item/"
	Item         = func(itemID int) string { return endpointItem + strconv.Itoa(itemID) }        // ITEM 		/WOW/ITEM/:ITEMID
	ItemSet      = func(setID int) string { return endpointItem + "set/" + strconv.Itoa(setID) } // ITEM SET 	/WOW/ITEM/SET/:SETID

	//	MOUNT
	Mount = wowAPI + "mount/" // MASTER LIST /WOW/MOUNT/

	//	PET
	PetMasterList = wowAPI + "pet/"                                                                            // MASTER LIST 	/WOW/PET/
	PetAbilities  = func(abilityID int) string { return PetMasterList + "ability/" + strconv.Itoa(abilityID) } // ABILITIES 		/WOW/PET/ABILITY/:ABILITYID
	PetSpecies    = func(speciesID int) string { return PetMasterList + "species/" + strconv.Itoa(speciesID) } // SPECIES 		/WOW/PET/SPECIES/:SPECIESID
	PetStats      = func(speciesID int) string { return PetMasterList + "stats/" + strconv.Itoa(speciesID) }   // STATS 			/WOW/PET/STATS/:SPECIESID

	//	PVP
	endpointPvp  = wowAPI + "leaderboard/"
	Leaderboards = func(bracket string) string { return endpointPvp + bracket } // LEADERBOARDS /WOW/LEADERBOARD/:BRACKET

	//	QUEST
	endpointQuest = wowAPI + "quest/"
	Quest         = func(questID int) string { return endpointQuest + strconv.Itoa(questID) } // QUEST /WOW/QUEST/:QUESTID

	//	REALM STATUS
	endpointRealm = wowAPI + "realm/"
	RealmStatus   = endpointRealm + "status" // REALM STATUS /WOW/REALM/STATUS

	//	RECIPE
	endpointRecipe = wowAPI + "recipe/"
	Recipe         = func(recipeID int) string { return endpointRecipe + strconv.Itoa(recipeID) } // RECIPE /WOW/RECIPE/:RECIPEID

	//	SPELL
	endpointSpell = wowAPI + "spell/"
	Spell         = func(spellID int) string { return endpointSpell + strconv.Itoa(spellID) } // SPELL /WOW/SPELL/:SPELLID

	//	ZONE
	ZoneMasterList = wowAPI + "zone/"                                                         // MASTER LIST 	/WOW/ZONE/
	Zone           = func(zoneID int) string { return ZoneMasterList + strconv.Itoa(zoneID) } // ZONE 		/WOW/ZONE/:ZONEID

	//	DATA RESOURCES
	endpointData          = wowAPI + "data/"
	endpointDataCharacter = endpointData + "character/"
	endpointDataGuild     = endpointData + "guild/"
	Battlegroups          = endpointData + "battlegroups/"         // BATTLEGROUPS 			/WOW/DATA/BATTLEGROUPS/
	CharacterRaces        = endpointDataCharacter + "races"        // CHARACTER RACES 		/WOW/DATA/CHARACTER/RACES
	CharacterClasses      = endpointDataCharacter + "classes"      // CHARACTER CLASSES 		/WOW/DATA/CHARACTER/CLASSES
	CharacterAchievements = endpointDataCharacter + "achievements" // CHARACTER ACHIEVEMENTS 	/WOW/DATA/CHARACTER/ACHIEVEMENTS
	GuildRewards          = endpointDataGuild + "rewards"          // GUILD REWARDS 			/WOW/DATA/GUILD/REWARDS
	GuildPerks            = endpointDataGuild + "perks"            // GUILD PERKS 			/WOW/DATA/GUILD/PERKS
	GuildAchievements     = endpointDataGuild + "achievements"     // GUILD ACHIEVEMENTS 		/WOW/DATA/GUILD/ACHIEVEMENTS
	ItemClasses           = endpointData + "item/classes"          // ITEM CLASSES 			/WOW/DATA/ITEM/CLASSES
	Talents               = endpointData + "talents"               // TALENTS 				/WOW/DATA/TALENTS
	PetTypes              = endpointData + "pet/types"             // PET TYPES 				/WOW/DATA/PET/TYPES

	//	WOW GAME DATA API
	//	CONNECTED REALM
	GetConnectedRealmIndex = "/connected-realm/"                 // GETCONNECTEDREALMINDEX 	/CONNECTED-REALM/
	GetConnectedRealm      = func(connectedRealmID int) string { // GETCONNECTEDREALM 		/CONNECTED-REALM/{CONNECTEDREALMID}
		return GetConnectedRealmIndex + strconv.Itoa(connectedRealmID)
	}

	//	MYTHIC KEYSTONE LEADERBOARD
	GetMythicLeaderboardIndex = func(connectedRealmID int) string { // GETMYTHICLEADERBOARDINDEX 	/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/
		return GetConnectedRealm(connectedRealmID) + "/mythic-leaderboard/"
	}
	GetMythicLeaderboard = func(connectedRealmID, dungeonID, period int) string { // GETMYTHICLEADERBOARD 	/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/{DUNGEONID}/PERIOD/{PERIOD}
		return GetMythicLeaderboardIndex(connectedRealmID) + strconv.Itoa(dungeonID) + "/period/" + strconv.Itoa(period)
	}

	//	MYTHIC CHALLENGE MODE
	GetMythicChallengeModeIndex = "/mythic-challenge-mode/" // GETMYTHICCHALLENGEMODEINDEX /MYTHIC-CHALLENGE-MODE/

	//	PLAYABLE CLASS
	GetPlayableClassesIndex = "/playable-class/"                                                                  // GETPLAYABLECLASSESINDEX /PLAYABLE-CLASS/
	GetPlayableClass        = func(classID int) string { return GetPlayableClassesIndex + strconv.Itoa(classID) } // GETPLAYABLECLASS 		 /PLAYABLE-CLASS/{CLASSID}

	//	PLAYABLE SPECIALIZATION
	GetPlayableSpecializationIndex = "/playable-specialization/"                                                              // GETPLAYABLESPECIALIZATIONINDEX /PLAYABLE-SPECIALIZATION/
	GetPlayableSpecialization      = func(specID int) string { return GetPlayableSpecializationIndex + strconv.Itoa(specID) } // GETPLAYABLESPECIALIZATION 		/PLAYABLE-SPECIALIZATION/{SPECID}

	//	REALM
	GetRealmIndex = "/realm/"                                                          // GETREALMINDEX 	/REALM/
	GetRealm      = func(realmslug string) string { return GetRealmIndex + realmslug } // GETREALM 		/REALM/{REALMSLUG}

	//	REGION
	GetRegionIndex = "/region/"                                                                   // GETREGIONINDEX 	/REGION/
	GetRegion      = func(regionID int) string { return GetRegionIndex + strconv.Itoa(regionID) } // GETREGION 		/REGION/{REGIONID}

	//	WOW TOKEN
	GetTokenIndex = "/token/" // GETTOKENINDEX /TOKEN/
)

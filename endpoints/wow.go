package endpoints

import (
	"strconv"
)

var (
	wow = API() + "wow/"

	//	PROFILE
	WowUser         = wow + "user/"
	WowOauthProfile = WowUser + "characters" // WOW OAUTH PROFILE /WOW/USER/CHARACTERS
	//	ACHIEVEMENT
	WowAchievement = func(id int) string { return wow + "achievement/" + strconv.Itoa(id) } // ACHIEVEMENT /WOW/ACHIEVEMENT/:ID

	//	AUCTION
	WowAuction           = wow + "auction/"
	WowAuctionDataStatus = func(realm string) string { return WowAuction + "data/" + realm } // AUCTION DATA STATUS /WOW/AUCTION/DATA/:REALM

	//	BOSS
	WowBossMaster = wow + "boss/"                                                   // MASTER LIST /WOW/BOSS/
	WowBoss       = func(id int) string { return WowBossMaster + strconv.Itoa(id) } // BOSS /WOW/BOSS/:BOSSID

	//	CHALLENGE MODE
	WowChallenge         = wow + "challenge/"
	WowRealmLeaderboard  = func(realm string) string { return WowChallenge + realm } // REALM LEADERBOARD /WOW/CHALLENGE/:REALM
	WowRegionLeaderboard = WowChallenge + "region"                                   // REGION LEADERBOARD /WOW/CHALLENGE/REGION

	//	CHARACTER PROFILE
	WowCharacter        = wow + "character/"
	WowCharacterProfile = func(realm, characterName string) string { return WowCharacter + realm + "/" + characterName } // CHARACTER PROFILE 	/WOW/CHARACTER/:REALM/:CHARACTERNAME
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
	WowGuild        = wow + "guild/"
	WowGuildProfile = func(realm, guildName string) string { return WowGuild + realm + "/" + guildName } // GUILD PROFILE 	/WOW/GUILD/:REALM/:GUILDNAME
	// MEMBERS 			/WOW/GUILD/:REALM/:GUILDNAME // implement the rest differently?
	// ACHIEVEMENTS 	/WOW/GUILD/:REALM/:GUILDNAME
	// NEWS 			/WOW/GUILD/:REALM/:GUILDNAME
	// CHALLENGE 		/WOW/GUILD/:REALM/:GUILDNAME

	//	ITEM
	endpointWowItem = wow + "item/"
	WowItem         = func(itemID int) string { return endpointWowItem + strconv.Itoa(itemID) }        // ITEM 		/WOW/ITEM/:ITEMID
	WowItemSet      = func(setID int) string { return endpointWowItem + "set/" + strconv.Itoa(setID) } // ITEM SET 	/WOW/ITEM/SET/:SETID

	//	MOUNT
	WowMount = wow + "mount/" // MASTER LIST /WOW/MOUNT/

	//	PET
	WowPet          = wow + "pet/"                                                                        // MASTER LIST 	/WOW/PET/
	WowPetAbilities = func(abilityID int) string { return WowPet + "ability/" + strconv.Itoa(abilityID) } // ABILITIES 		/WOW/PET/ABILITY/:ABILITYID
	WowPetSpecies   = func(speciesID int) string { return WowPet + "species/" + strconv.Itoa(speciesID) } // SPECIES 		/WOW/PET/SPECIES/:SPECIESID
	WowPetStats     = func(speciesID int) string { return WowPet + "stats/" + strconv.Itoa(speciesID) }   // STATS 			/WOW/PET/STATS/:SPECIESID

	//	PVP
	WowPvp          = wow + "leaderboard/"
	WowLeaderboards = func(bracket string) string { return WowPvp + bracket } // LEADERBOARDS /WOW/LEADERBOARD/:BRACKET

	//	QUEST
	endpointWowQuest = wow + "quest/"
	WowQuest         = func(questID int) string { return endpointWowQuest + strconv.Itoa(questID) } // QUEST /WOW/QUEST/:QUESTID

	//	REALM STATUS
	endpointWowRealm = wow + "realm/"
	WowRealmStatus   = endpointRealm + "status" // REALM STATUS /WOW/REALM/STATUS

	//	RECIPE
	endpointWowRecipe = wow + "recipe/"
	WowRecipe         = func(recipeID int) string { return endpointWowRecipe + strconv.Itoa(recipeID) } // RECIPE /WOW/RECIPE/:RECIPEID

	//	SPELL
	endpointWowSpell = wow + "spell/"
	WowSpell         = func(spellID int) string { return endpointWowSpell + strconv.Itoa(spellID) } // SPELL /WOW/SPELL/:SPELLID

	//	ZONE
	WowZoneList = wow + "zone/"                                                         // MASTER LIST 	/WOW/ZONE/
	WowZone     = func(zoneID int) string { return WowZoneList + strconv.Itoa(zoneID) } // ZONE 		/WOW/ZONE/:ZONEID

	//	DATA RESOURCES
	endpointWowData          = wow + "data/"
	endpointWowDataCharacter = endpointWowData + "character/"
	endpointWowDataGuild     = endpointWowData + "guild/"
	WowBattlegroups          = endpointWowData + "battlegroups/"         // BATTLEGROUPS 			/WOW/DATA/BATTLEGROUPS/
	WowCharacterRaces        = endpointWowDataCharacter + "races"        // CHARACTER RACES 		/WOW/DATA/CHARACTER/RACES
	WowCharacterClasses      = endpointWowDataCharacter + "classes"      // CHARACTER CLASSES 		/WOW/DATA/CHARACTER/CLASSES
	WowCharacterAchievements = endpointWowDataCharacter + "achievements" // CHARACTER ACHIEVEMENTS 	/WOW/DATA/CHARACTER/ACHIEVEMENTS
	WowGuildRewards          = endpointWowDataGuild + "rewards"          // GUILD REWARDS 			/WOW/DATA/GUILD/REWARDS
	WowGuildPerks            = endpointWowDataGuild + "perks"            // GUILD PERKS 			/WOW/DATA/GUILD/PERKS
	WowGuildAchievements     = endpointWowDataGuild + "achievements"     // GUILD ACHIEVEMENTS 		/WOW/DATA/GUILD/ACHIEVEMENTS
	WowItemClasses           = endpointWowData + "item/classes"          // ITEM CLASSES 			/WOW/DATA/ITEM/CLASSES
	WowTalents               = endpointWowData + "talents"               // TALENTS 				/WOW/DATA/TALENTS
	WowPetTypes              = WowData + "pet/types"                     // PET TYPES 				/WOW/DATA/PET/TYPES

	//	WOW GAME DATA API
	//	CONNECTED REALM
	GetConnectedRealmIndex = "/connected-realm/" // GETCONNECTEDREALMINDEX 	/CONNECTED-REALM/
	// GETCONNECTEDREALM 		/CONNECTED-REALM/{CONNECTEDREALMID}
	GetConnectedRealm = func(connectedRealmID int) string {
		return GetConnectedRealmIndex + strconv.Itoa(connectedRealmID)
	}

	//	MYTHIC KEYSTONE LEADERBOARD
	// GETMYTHICLEADERBOARDINDEX 	/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/
	GetMythicLeaderboardIndex = func(connectedRealmID int) string {
		return GetConnectedRealm(connectedRealmID) + "/mythic-leaderboard/"
	}
	// GETMYTHICLEADERBOARD 		/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/{DUNGEONID}/PERIOD/{PERIOD}
	GetMythicLeaderboard = func(connectedRealmID, dungeonID, period int) string {
		return GetMythicLeaderboardIndex(connectedRealmID) + strconv.Itoa(dungeonID) + "/period/" + strconv.Itoa(period)
	}

	//	MYTHIC CHALLENGE MODE
	GetMythicChallengeModeIndex = "/mythic-challenge-mode/" // GETMYTHICCHALLENGEMODEINDEX /MYTHIC-CHALLENGE-MODE/

	//	PLAYABLE CLASS
	GetPlayableClassesIndex = "/playable-class/"                                                                  // GETPLAYABLECLASSESINDEX /PLAYABLE-CLASS/
	GetPlayableClass        = func(classID int) string { return GetPlayableClassesIndex + strconv.Itoa(classID) } // GETPLAYABLECLASS 		/PLAYABLE-CLASS/{CLASSID}

	//	PLAYABLE SPECIALIZATION
	GetPlayableSpecializationIndex = "/playable-specialization/"                                                              // GETPLAYABLESPECIALIZATIONINDEX 	/PLAYABLE-SPECIALIZATION/
	GetPlayableSpecialization      = func(specID int) string { return GetPlayableSpecializationIndex + strconv.Itoa(specID) } // GETPLAYABLESPECIALIZATION 		/PLAYABLE-SPECIALIZATION/{SPECID}

	//	REALM
	GetRealmIndex = "/realm/"                                                          // GETREALMINDEX 	/REALM/
	GetRealm      = func(realmslug string) string { return GetRealmIndex + realmslug } // GETREALM 		/REALM/{REALMSLUG}

	//	REGION
	GetRegionIndex = "/region/"                                                                   // GETREGIONINDEX 	/REGION/
	GetRegion      = func(regionID int) string { return GetRegionIndex + strconv.Itoa(regionID) } // GETREGION 		/REGION/{REGIONID}

	//	WOW TOKEN
	GetWowTokenIndex = "/token/" // GETTOKENINDEX /TOKEN/
)

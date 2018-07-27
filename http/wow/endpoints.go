package wow

import (
	"strconv"

	"github.com/munsy/battlenet/regions"
)

// r regions.Region,
var (
	endpointWow = func(r regions.Region) string { return r.API() + "wow/" }

	//	ACHIEVEMENT
	endpointAchievement = func(r regions.Region, id int) string { return endpointWow(r) + "achievement/" + strconv.Itoa(id) } // ACHIEVEMENT /WOW/ACHIEVEMENT/:ID

	//	AUCTION
	endpointAuction           = func(r regions.Region) string { return endpointWow(r) + "auction/" }
	endpointAuctionDataStatus = func(r regions.Region, realm string) string { return endpointAuction(r) + "data/" + realm } // AUCTION DATA STATUS /WOW/AUCTION/DATA/:REALM

	//	BOSS
	endpointBossMasterList = func(r regions.Region) string { return endpointWow(r) + "boss/" }                             // MASTER LIST /WOW/BOSS/
	endpointBossInfo       = func(r regions.Region, id int) string { return endpointBossMasterList(r) + strconv.Itoa(id) } // BOSS /WOW/BOSS/:BOSSID

	//	CHARACTER PROFILE
	endpointCharacter        = func(r regions.Region) string { return endpointWow(r) + "character/" }
	endpointCharacterProfile = func(r regions.Region, realm, characterName string) string {
		return endpointCharacter(r) + realm + "/" + characterName
	} // CHARACTER PROFILE 	/WOW/CHARACTER/:REALM/:CHARACTERNAME
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
	endpointGuild        = func(r regions.Region) string { return endpointWow(r) + "guild/" }
	endpointGuildProfile = func(r regions.Region, realm, guildName string) string {
		return endpointGuild(r) + realm + "/" + guildName
	} // GUILD PROFILE 	/WOW/GUILD/:REALM/:GUILDNAME
	// MEMBERS 			/WOW/GUILD/:REALM/:GUILDNAME // implement the rest differently?
	// ACHIEVEMENTS 	/WOW/GUILD/:REALM/:GUILDNAME
	// NEWS 			/WOW/GUILD/:REALM/:GUILDNAME
	// CHALLENGE 		/WOW/GUILD/:REALM/:GUILDNAME

	//	ITEM
	endpointItem    = func(r regions.Region) string { return endpointWow(r) + "item/" }
	endpointItemID  = func(r regions.Region, itemID int) string { return endpointItem(r) + strconv.Itoa(itemID) }        // ITEM 		/WOW/ITEM/:ITEMID
	endpointItemSet = func(r regions.Region, setID int) string { return endpointItem(r) + "set/" + strconv.Itoa(setID) } // ITEM SET 	/WOW/ITEM/SET/:SETID

	//	MOUNT
	endpointMount = func(r regions.Region) string { return endpointWow(r) + "mount/" } // MASTER LIST /WOW/MOUNT/

	//	PET
	endpointPetMasterList = func(r regions.Region) string { return endpointWow(r) + "pet/" } // MASTER LIST 	/WOW/PET/
	endpointPetAbilities  = func(r regions.Region, abilityID int) string {
		return endpointPetMasterList(r) + "ability/" + strconv.Itoa(abilityID)
	} // ABILITIES 		/WOW/PET/ABILITY/:ABILITYID
	endpointPetSpecies = func(r regions.Region, speciesID int) string {
		return endpointPetMasterList(r) + "species/" + strconv.Itoa(speciesID)
	} // SPECIES 		/WOW/PET/SPECIES/:SPECIESID
	endpointPetStats = func(r regions.Region, speciesID int) string {
		return endpointPetMasterList(r) + "stats/" + strconv.Itoa(speciesID)
	} // STATS 			/WOW/PET/STATS/:SPECIESID

	//	PVP
	endpointPvp          = func(r regions.Region) string { return endpointWow(r) + "leaderboard/" }
	endpointLeaderboards = func(r regions.Region, bracket string) string { return endpointPvp(r) + bracket } // LEADERBOARDS /WOW/LEADERBOARD/:BRACKET

	//	QUEST
	endpointQuest   = func(r regions.Region) string { return endpointWow(r) + "quest/" }
	endpointQuestID = func(r regions.Region, questID int) string { return endpointQuest(r) + strconv.Itoa(questID) } // QUEST /WOW/QUEST/:QUESTID

	//	REALM STATUS
	endpointRealm       = func(r regions.Region) string { return endpointWow(r) + "realm/" }
	endpointRealmStatus = func(r regions.Region) string { return endpointRealm(r) + "status" } // REALM STATUS /WOW/REALM/STATUS

	//	RECIPE
	endpointRecipe   = func(r regions.Region) string { return endpointWow(r) + "recipe/" }
	endpointRecipeID = func(r regions.Region, recipeID int) string { return endpointRecipe(r) + strconv.Itoa(recipeID) } // RECIPE /WOW/RECIPE/:RECIPEID

	//	SPELL
	endpointSpell   = func(r regions.Region) string { return endpointWow(r) + "spell/" }
	endpointSpellID = func(r regions.Region, spellID int) string { return endpointSpell(r) + strconv.Itoa(spellID) } // SPELL /WOW/SPELL/:SPELLID

	//	ZONE
	endpointZoneMasterList = func(r regions.Region) string { return endpointWow(r) + "zone/" }                                     // MASTER LIST 	/WOW/ZONE/
	endpointZone           = func(r regions.Region, zoneID int) string { return endpointZoneMasterList(r) + strconv.Itoa(zoneID) } // ZONE 		/WOW/ZONE/:ZONEID

	//	DATA RESOURCES
	endpointData                  = func(r regions.Region) string { return endpointWow(r) + "data/" }
	endpointDataCharacter         = func(r regions.Region) string { return endpointData(r) + "character/" }
	endpointDataGuild             = func(r regions.Region) string { return endpointData(r) + "guild/" }
	endpointBattlegroups          = func(r regions.Region) string { return endpointData(r) + "battlegroups/" }         // BATTLEGROUPS 			/WOW/DATA/BATTLEGROUPS/
	endpointCharacterRaces        = func(r regions.Region) string { return endpointDataCharacter(r) + "races" }        // CHARACTER RACES 		/WOW/DATA/CHARACTER/RACES
	endpointCharacterClasses      = func(r regions.Region) string { return endpointDataCharacter(r) + "classes" }      // CHARACTER CLASSES 		/WOW/DATA/CHARACTER/CLASSES
	endpointCharacterAchievements = func(r regions.Region) string { return endpointDataCharacter(r) + "achievements" } // CHARACTER ACHIEVEMENTS 	/WOW/DATA/CHARACTER/ACHIEVEMENTS
	endpointGuildRewards          = func(r regions.Region) string { return endpointDataGuild(r) + "rewards" }          // GUILD REWARDS 			/WOW/DATA/GUILD/REWARDS
	endpointGuildPerks            = func(r regions.Region) string { return endpointDataGuild(r) + "perks" }            // GUILD PERKS 			/WOW/DATA/GUILD/PERKS
	endpointGuildAchievements     = func(r regions.Region) string { return endpointDataGuild(r) + "achievements" }     // GUILD ACHIEVEMENTS 		/WOW/DATA/GUILD/ACHIEVEMENTS
	endpointItemClasses           = func(r regions.Region) string { return endpointData(r) + "item/classes" }          // ITEM CLASSES 			/WOW/DATA/ITEM/CLASSES
	endpointTalents               = func(r regions.Region) string { return endpointData(r) + "talents" }               // TALENTS 				/WOW/DATA/TALENTS
	endpointPetTypes              = func(r regions.Region) string { return endpointData(r) + "pet/types" }             // PET TYPES 				/WOW/DATA/PET/TYPES

	//	WOW GAME DATA API
	//	CONNECTED REALM
	endpointGetConnectedRealmIndex = func(r regions.Region) string { return r.API() + "/connected-realm/" } // GETCONNECTEDREALMINDEX /CONNECTED-REALM/
	endpointGetConnectedRealm      = func(r regions.Region, connectedRealmID int) string {                  // GETCONNECTEDREALM 		/CONNECTED-REALM/{CONNECTEDREALMID}
		return endpointGetConnectedRealmIndex(r) + strconv.Itoa(connectedRealmID)
	}

	//	MYTHIC KEYSTONE LEADERBOARD
	endpointGetMythicLeaderboardIndex = func(r regions.Region, connectedRealmID int) string { // GETMYTHICLEADERBOARDINDEX 	/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/
		return endpointGetConnectedRealm(r, connectedRealmID) + "/mythic-leaderboard/"
	}
	endpointGetMythicLeaderboard = func(r regions.Region, connectedRealmID, dungeonID, period int) string { // GETMYTHICLEADERBOARD 	/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/{DUNGEONID}/PERIOD/{PERIOD}
		return endpointGetMythicLeaderboardIndex(r, connectedRealmID) + strconv.Itoa(dungeonID) + "/period/" + strconv.Itoa(period)
	}

	//	MYTHIC CHALLENGE MODE
	endpointGetMythicChallengeModeIndex = func(r regions.Region) string { return r.API() + "/mythic-challenge-mode/" } // GETMYTHICCHALLENGEMODEINDEX /MYTHIC-CHALLENGE-MODE/

	//	PLAYABLE CLASS
	endpointGetPlayableClassesIndex = func(r regions.Region) string { return r.API() + "/playable-class/" } // GETPLAYABLECLASSESINDEX /PLAYABLE-CLASS/
	endpointGetPlayableClass        = func(r regions.Region, classID int) string {
		return endpointGetPlayableClassesIndex(r) + strconv.Itoa(classID)
	} // GETPLAYABLECLASS 		 /PLAYABLE-CLASS/{CLASSID}

	//	PLAYABLE SPECIALIZATION
	endpointGetPlayableSpecializationIndex = func(r regions.Region) string { return r.API() + "/playable-specialization/" } // GETPLAYABLESPECIALIZATIONINDEX /PLAYABLE-SPECIALIZATION/
	endpointGetPlayableSpecialization      = func(r regions.Region, specID int) string {
		return endpointGetPlayableSpecializationIndex(r) + strconv.Itoa(specID)
	} // GETPLAYABLESPECIALIZATION 		/PLAYABLE-SPECIALIZATION/{SPECID}

	//	REALM
	endpointGetRealmIndex = func(r regions.Region) string { return r.API() + "/realm/" }                                    // GETREALMINDEX 	/REALM/
	endpointGetRealm      = func(r regions.Region, realmslug string) string { return endpointGetRealmIndex(r) + realmslug } // GETREALM 		/REALM/{REALMSLUG}

	//	REGION
	endpointGetRegionIndex = func(r regions.Region) string { return r.API() + "/region/" } // GETREGIONINDEX 	/REGION/
	endpointGetRegion      = func(r regions.Region, regionID int) string {
		return endpointGetRegionIndex(r) + strconv.Itoa(regionID)
	} // GETREGION 		/REGION/{REGIONID}

	//	WOW TOKEN
	endpointGetTokenIndex = func(r regions.Region) string { return r.API() + "/token/" } // GETTOKENINDEX /TOKEN/
)

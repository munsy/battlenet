package endpoints

import (
	"strconv"
)

var (
	//	WOW API
	//	------------------------------------------------------------
	EndpointWow = EndpointAPI() + "wow/"
	//	PROFILE
	EndpointWowUser = EndpointWow + "user/"
	//[x]	WOW OAUTH PROFILE /WOW/USER/CHARACTERS\
	EndpointWowOauthProfile = EndpointWowUser + "characters"
	//	ACHIEVEMENT
	//[x]	ACHIEVEMENT /WOW/ACHIEVEMENT/:ID
	EndpointWowAchievement = func(id int) string { return EndpointWow + "achievement/" + strconv.Itoa(id) }

	//	AUCTION
	EndpointWowAuction = EndpointWow + "auction/"
	//[x]	AUCTION DATA STATUS /WOW/AUCTION/DATA/:REALM
	EndpointWowAuctionDataStatus = func(realm string) string { return EndpointWowAuction + "data/" + realm }

	//	BOSS
	//[x]	MASTER LIST /WOW/BOSS/
	EndpointWowBossMaster = EndpointWow + "boss/"
	//[x]	BOSS /WOW/BOSS/:BOSSID
	EndpointWowBoss = func(id int) string { return EndpointWowBossMaster + strconv.Itoa(id) }

	//	CHALLENGE MODE
	EndpointWowChallenge = EndpointWow + "challenge/"
	//[x]	REALM LEADERBOARD /WOW/CHALLENGE/:REALM
	EndpointWowRealmLeaderboard = func(realm string) string { return EndpointWowChallenge + realm }
	//[x]	REGION LEADERBOARD /WOW/CHALLENGE/REGION
	EndpointWowRegionLeaderboard = EndpointWowChallenge + "region"

	//	CHARACTER PROFILE
	EndpointWowCharacter = EndpointWow + "character/"
	//[x]	CHARACTER PROFILE 	/WOW/CHARACTER/:REALM/:CHARACTERNAME // Implement the rest differently?
	EndpointWowCharacterProfile = func(realm, characterName string) string { return EndpointWowCharacter + realm + "/" + characterName }
	//[ ]	ACHIEVEMENTS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	APPEARANCE 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	FEED 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	GUILD 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	HUNTER PETS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	ITEMS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	MOUNTS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	PETS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	PET SLOTS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	PROFESSIONS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	PROGRESSION 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	PVP 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	QUESTS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	REPUTATION 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	STATISTICS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	STATS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	TALENTS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	TITLES 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
	//[ ]	AUDIT 				/WOW/CHARACTER/:REALM/:CHARACTERNAME

	//	GUILD PROFILE
	EndpointWowGuild = EndpointWow + "guild/"
	//[x]	GUILD PROFILE 	/WOW/GUILD/:REALM/:GUILDNAME // implement the rest differently?
	EndpointWowGuildProfile = func(realm, guildName string) string { return EndpointWowGuild + realm + "/" + guildName }
	//[ ]	MEMBERS 		/WOW/GUILD/:REALM/:GUILDNAME
	//[ ]	ACHIEVEMENTS 	/WOW/GUILD/:REALM/:GUILDNAME
	//[ ]	NEWS 			/WOW/GUILD/:REALM/:GUILDNAME
	//[ ]	CHALLENGE 		/WOW/GUILD/:REALM/:GUILDNAME

	//	ITEM
	endpointWowItem = EndpointWow + "item/"
	//[x]	ITEM 		/WOW/ITEM/:ITEMID
	EndpointWowItem = func(itemID int) string { return endpointWowItem + strconv.Itoa(itemID) }
	//[x]	ITEM SET 	/WOW/ITEM/SET/:SETID
	EndpointWowItemSet = func(setID int) string { return endpointWowItem + "set/" + strconv.Itoa(setID) }

	//	MOUNT
	//[x]	MASTER LIST /WOW/MOUNT/
	EndpointWowMount = EndpointWow + "mount/"

	//	PET
	//[x]	MASTER LIST /WOW/PET/
	EndpointWowPet = EndpointWow + "pet/"
	//[x]	ABILITIES 	/WOW/PET/ABILITY/:ABILITYID
	EndpointWowPetAbilities = func(abilityID int) string { return EndpointWowPet + "ability/" + strconv.Itoa(abilityID) }
	//[x]	SPECIES 	/WOW/PET/SPECIES/:SPECIESID
	EndpointWowPetSpecies = func(speciesID int) string { return EndpointWowPet + "species/" + strconv.Itoa(speciesID) }
	//[x]	STATS 		/WOW/PET/STATS/:SPECIESID
	EndpointWowPetStats = func(speciesID int) string { return EndpointWowPet + "stats/" + strconv.Itoa(speciesID) }

	//	PVP
	EndpointWowPvp = EndpointWow + "leaderboard/"
	//[x]	LEADERBOARDS /WOW/LEADERBOARD/:BRACKET
	EndpointWowLeaderboards = func(bracket string) string { return EndpointWowPvp + bracket }

	//	QUEST
	endpointWowQuest = EndpointWow + "quest/"
	//[x]	QUEST /WOW/QUEST/:QUESTID
	EndpointWowQuest = func(questID int) string { return endpointWowQuest + strconv.Itoa(questID) }

	//	REALM STATUS
	endpointWowRealm = EndpointWow + "realm/"
	//[x]	REALM STATUS /WOW/REALM/STATUS
	EndpointWowRealmStatus = endpointRealm + "status"

	//	RECIPE
	endpointWowRecipe = EndpointWow + "recipe/"
	//[x]	RECIPE /WOW/RECIPE/:RECIPEID
	EndpointWowRecipe = func(recipeID int) string { return endpointWowRecipe + strconv.Itoa(recipeID) }

	//	SPELL
	endpointWowSpell = EndpointWow + "spell/"
	//[x]	SPELL /WOW/SPELL/:SPELLID
	EndpointWowSpell = func(spellID int) string { return endpointWowSpell + strconv.Itoa(spellID) }

	//	ZONE
	//[x]	MASTER LIST /WOW/ZONE/
	EndpointWowZoneList = EndpointWow + "zone/"
	//[x]	ZONE 		/WOW/ZONE/:ZONEID
	EndpointWowZone = func(zoneID int) string { return EndpointWowZoneList + strconv.Itoa(zoneID) }

	//	DATA RESOURCES
	endpointWowData          = EndpointWow + "data/"
	endpointWowDataCharacter = endpointWowData + "character/"
	endpointWowDataGuild     = endpointWowData + "guild/"
	//[x]	BATTLEGROUPS 			/WOW/DATA/BATTLEGROUPS/
	EndpointWowBattlegroups = endpointWowData + "battlegroups/"
	//[x]	CHARACTER RACES 		/WOW/DATA/CHARACTER/RACES
	EndpointWowCharacterRaces = endpointWowDataCharacter + "races"
	//[x]	CHARACTER CLASSES 		/WOW/DATA/CHARACTER/CLASSES
	EndpointWowCharacterClasses = endpointWowDataCharacter + "classes"
	//[x]	CHARACTER ACHIEVEMENTS 	/WOW/DATA/CHARACTER/ACHIEVEMENTS
	EndpointWowCharacterAchievements = endpointWowDataCharacter + "achievements"
	//[x]	GUILD REWARDS 			/WOW/DATA/GUILD/REWARDS
	EndpointWowGuildRewards = endpointWowDataGuild + "rewards"
	//[x]	GUILD PERKS 			/WOW/DATA/GUILD/PERKS
	EndpointWowGuildPerks = endpointWowDataGuild + "perks"
	//[x]	GUILD ACHIEVEMENTS 		/WOW/DATA/GUILD/ACHIEVEMENTS
	EndpointWowGuildAchievements = endpointWowDataGuild + "achievements"
	//[x]	ITEM CLASSES 			/WOW/DATA/ITEM/CLASSES
	EndpointWowItemClasses = endpointWowData + "item/classes"
	//[x]	TALENTS 				/WOW/DATA/TALENTS
	EndpointWowTalents = endpointWowData + "talents"
	//[x]	PET TYPES 				/WOW/DATA/PET/TYPES
	EndpointWowPetTypes = EndpointWowData + "pet/types"

	//	WOW GAME DATA API
	//	--------------------------------------------------------------------
	//	CONNECTED REALM
	EndpointGet = ""
	//[ ]	GETCONNECTEDREALMINDEX 	/CONNECTED-REALM/
	EndpointGetConnectedRealmIndex = "/connected-realm/"
	//[ ]	GETCONNECTEDREALM 		/CONNECTED-REALM/{CONNECTEDREALMID}
	EndpointGetConnectedRealm = func(connectedRealmID int) string {
		return EndpointGetConnectedRealmIndex + strconv.Itoa(connectedRealmID)
	}

	//	MYTHIC KEYSTONE LEADERBOARD
	//[ ]	GETMYTHICLEADERBOARDINDEX 	/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/
	EndpointGetMythicLeaderboardIndex = func(connectedRealmID int) string {
		return EndpointGetConnectedRealm(connectedRealmID) + "/mythic-leaderboard/"
	}
	//[ ]	GETMYTHICLEADERBOARD 		/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/{DUNGEONID}/PERIOD/{PERIOD}
	EndpointGetMythicLeaderboard = func(connectedRealmID, dungeonID, period int) string {
		return EndpointGetMythicLeaderboardIndex(connectedRealmID) + strconv.Itoa(dungeonID) + "/period/" + strconv.Itoa(period)
	}

	//	MYTHIC CHALLENGE MODE
	//[ ]	GETMYTHICCHALLENGEMODEINDEX /MYTHIC-CHALLENGE-MODE/
	EndpointGetMythicChallengeModeIndex = "/mythic-challenge-mode/"

	//	PLAYABLE CLASS
	//[ ]	GETPLAYABLECLASSESINDEX /PLAYABLE-CLASS/
	EndpointGetPlayableClassesIndex = "/playable-class/"
	//[ ]	GETPLAYABLECLASS 		/PLAYABLE-CLASS/{CLASSID}
	EndpointGetPlayableClass = func(classID int) string { return EndpointGetPlayableClassesIndex + strconv.Itoa(classID) }

	//	PLAYABLE SPECIALIZATION
	//[ ]	GETPLAYABLESPECIALIZATIONINDEX 	/PLAYABLE-SPECIALIZATION/
	EndpointGetPlayableSpecializationIndex = "/playable-specialization/"
	//[ ]	GETPLAYABLESPECIALIZATION 		/PLAYABLE-SPECIALIZATION/{SPECID}
	EndpointGetPlayableSpecialization = func(specID int) string { return EndpointGetPlayableSpecializationIndex + strconv.Itoa(specID) }

	//	REALM
	//[ ]	GETREALMINDEX 	/REALM/
	EndpointGetRealmIndex = "/realm/"
	//[ ]	GETREALM 		/REALM/{REALMSLUG}
	EndpointGetRealm = func(realmslug string) string { return EndpointGetRealmIndex + realmslug }

	//	REGION
	//[ ]	GETREGIONINDEX 	/REGION/
	EndpointGetRegionIndex = "/region/"
	//[ ]	GETREGION 		/REGION/{REGIONID}
	EndpointGetRegion = func(regionID int) string { return EndpointGetRegionIndex + strconv.Itoa(regionID) }

	//	WOW TOKEN
	//[ ]	GETTOKENINDEX /TOKEN/
	EndpointGetWowTokenIndex = "/token/"
)

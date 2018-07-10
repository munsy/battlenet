package battlenet

import (
	"errors"
	"strconv"
	"strings"
)

var (
	EndpointOauth2 = func(endpoint string) string {
		if "token" != endpoint && "authorize" != endpoint {
			panic(errors.New("Can't resolve endpoint."))
		}

		r := strings.ToLower(region)

		switch r {
		case "us":
		case "eu":
		case "kr":
		case "tw":
		case "sea":
			return "https://" + r + ".battle.net/oauth/" + endpoint
		case "cn":
			return "https://www.battlenet.com.cn/oauth/" + endpoint
		default:
			return "https://us.battle.net/oauth/" + endpoint
		}
		return "https://us.battle.net/oauth/" + endpoint
	}

	EndpointAPI = func() string {
		r := strings.ToLower(region)
		switch r {
		case "us":
		case "eu":
		case "kr":
		case "tw":
		case "sea":
			return "https://" + r + ".api.battle.net/"
		case "cn":
			return "https://api.battlenet.com.cn/"
		default:
			return "https://us.api.battle.net/"
		}
		return "https://us.api.battle.net/"
	}

	EndpointAuthURL  = EndpointOauth2("authorize")
	EndpointTokenURL = EndpointOauth2("token")

	// Community Oauth Profile API
	EndpointAccount = EndpointAPI() + "account/"
	EndpointUser    = EndpointAccount + "user"

	/*
		COMMUNITY OAUTH PROFILE API
		------------------------------------------------------------

		ACCOUNT
		USER /ACCOUNT/USER // done

		PROFILE
		SC2 OAUTH PROFILE /SC2/PROFILE/USER
		WOW OAUTH PROFILE /WOW/USER/CHARACTERS

	*/

	EndpointWow                  = EndpointAPI() + "wow/"
	EndpointWowUser              = EndpointWow + "user/"
	EndpointWowOauthProfile      = EndpointWowUser + "characters"
	EndpointWowAchievement       = func(id uint64) string { return EndpointWow + "achievement/" + strconv.Itoa(id) }
	EndpointWowAuction           = EndpointWow + "auction/"
	EndpointWowAuctionDataStatus = func(realm string) { return EndpointWowAuction + "data/" + realm }
	EndpointWowBossMaster        = EndpointWow + "boss/"
	EndpointWowBoss              = func(id int) { return EndpointWowBossMaster + strconv.Itoa(id) }
	EndpointWowChallenge         = EndpointWow + "challenge/"
	EndpointWowRealmLeaderboard  = func(realm string) { return EndpointWowChallenge + realm }
	EndpointWowRegionLeaderboard = EndpointWowChallenge + "region"

	/*
			WOW API
			------------------------------------------------------------
			PROFILE
		[x]	WOW OAUTH PROFILE /WOW/USER/CHARACTERS

			ACHIEVEMENT
		[x]	ACHIEVEMENT /WOW/ACHIEVEMENT/:ID

			AUCTION
		[x]	AUCTION DATA STATUS /WOW/AUCTION/DATA/:REALM

			BOSS
		[x]	MASTER LIST /WOW/BOSS/
		[x]	BOSS /WOW/BOSS/:BOSSID

			CHALLENGE MODE
		[x]	REALM LEADERBOARD /WOW/CHALLENGE/:REALM
		[x]	REGION LEADERBOARD /WOW/CHALLENGE/REGION

			CHARACTER PROFILE
		[ ]	CHARACTER PROFILE 	/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	ACHIEVEMENTS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	APPEARANCE 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	FEED 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	GUILD 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	HUNTER PETS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	ITEMS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	MOUNTS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	PETS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	PET SLOTS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	PROFESSIONS 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	PROGRESSION 		/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	PVP 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	QUESTS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	REPUTATION 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	STATISTICS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	STATS 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	TALENTS 			/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	TITLES 				/WOW/CHARACTER/:REALM/:CHARACTERNAME
		[ ]	AUDIT 				/WOW/CHARACTER/:REALM/:CHARACTERNAME

			GUILD PROFILE
		[ ]	GUILD PROFILE 	/WOW/GUILD/:REALM/:GUILDNAME
		[ ]	MEMBERS 		/WOW/GUILD/:REALM/:GUILDNAME
		[ ]	ACHIEVEMENTS 	/WOW/GUILD/:REALM/:GUILDNAME
		[ ]	NEWS 			/WOW/GUILD/:REALM/:GUILDNAME
		[ ]	CHALLENGE 		/WOW/GUILD/:REALM/:GUILDNAME

			ITEM
		[ ]	ITEM 		/WOW/ITEM/:ITEMID
		[ ]	ITEM SET 	/WOW/ITEM/SET/:SETID

			MOUNT
		[ ]	MASTER LIST /WOW/MOUNT/

			PET
		[ ]	MASTER LIST /WOW/PET/
		[ ]	ABILITIES 	/WOW/PET/ABILITY/:ABILITYID
		[ ]	SPECIES 	/WOW/PET/SPECIES/:SPECIESID
		[ ]	STATS 		/WOW/PET/STATS/:SPECIESID

			PVP
		[ ]	LEADERBOARDS /WOW/LEADERBOARD/:BRACKET

			QUEST
		[ ]	QUEST /WOW/QUEST/:QUESTID

			REALM STATUS
		[ ]	REALM STATUS /WOW/REALM/STATUS

			RECIPE
		[ ]	RECIPE /WOW/RECIPE/:RECIPEID

			SPELL
		[ ]	SPELL /WOW/SPELL/:SPELLID

			ZONE
		[ ]	MASTER LIST /WOW/ZONE/
		[ ]	ZONE 		/WOW/ZONE/:ZONEID

			DATA RESOURCES
		[ ]	BATTLEGROUPS 			/WOW/DATA/BATTLEGROUPS/
		[ ]	CHARACTER RACES 		/WOW/DATA/CHARACTER/RACES
		[ ]	CHARACTER CLASSES 		/WOW/DATA/CHARACTER/CLASSES
		[ ]	CHARACTER ACHIEVEMENTS 	/WOW/DATA/CHARACTER/ACHIEVEMENTS
		[ ]	GUILD REWARDS 			/WOW/DATA/GUILD/REWARDS
		[ ]	GUILD PERKS 			/WOW/DATA/GUILD/PERKS
		[ ]	GUILD ACHIEVEMENTS 		/WOW/DATA/GUILD/ACHIEVEMENTS
		[ ]	ITEM CLASSES 			/WOW/DATA/ITEM/CLASSES
		[ ]	TALENTS 				/WOW/DATA/TALENTS
		[ ]	PET TYPES 				/WOW/DATA/PET/TYPES

	*/

	/*
			SC2 API
			--------------------------------------------------------------------
			PROFILE
		[ ]	PROFILE 		/SC2/PROFILE/:ID/:REGION/:NAME/
		[ ]	LADDERS 		/SC2/PROFILE/:ID/:REGION/:NAME/LADDERS
		[ ]	MATCH HISTORY 	/SC2/PROFILE/:ID/:REGION/:NAME/MATCHES

			LADDER
		[ ]	LADDER /SC2/LADDER/:ID

			DATA RESOURCES
		[ ]	ACHIEVEMENTS 	/SC2/DATA/ACHIEVEMENTS
		[ ]	REWARDS 		/SC2/DATA/REWARDS

	*/

	/*
			WOW GAME DATA API
			--------------------------------------------------------------------
			CONNECTED REALM
		[ ]	GETCONNECTEDREALMINDEX 	/CONNECTED-REALM/
		[ ]	GETCONNECTEDREALM 		/CONNECTED-REALM/{CONNECTEDREALMID}

			MYTHIC KEYSTONE LEADERBOARD
		[ ]	GETMYTHICLEADERBOARDINDEX 	/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/
		[ ]	GETMYTHICLEADERBOARD 		/CONNECTED-REALM/{CONNECTEDREALMID}/MYTHIC-LEADERBOARD/{DUNGEONID}/PERIOD/{PERIOD}

			MYTHIC CHALLENGE MODE
		[ ]	GETMYTHICCHALLENGEMODEINDEX /MYTHIC-CHALLENGE-MODE/

			PLAYABLE CLASS
		[ ]	GETPLAYABLECLASSESINDEX /PLAYABLE-CLASS/
		[ ]	GETPLAYABLECLASS 		/PLAYABLE-CLASS/{CLASSID}

			PLAYABLE SPECIALIZATION
		[ ]	GETPLAYABLESPECIALIZATIONINDEX 	/PLAYABLE-SPECIALIZATION/
		[ ]	GETPLAYABLESPECIALIZATION 		/PLAYABLE-SPECIALIZATION/{SPECID}

			REALM
		[ ]	GETREALMINDEX 	/REALM/
		[ ]	GETREALM 		/REALM/{REALMSLUG}

			REGION
		[ ]	GETREGIONINDEX 	/REGION/
		[ ]	GETREGION 		/REGION/{REGIONID}

			WOW TOKEN
		[ ]	GETTOKENINDEX /TOKEN/

	*/

	/*
			D3 API
			--------------------------------------------------------------------
			D3 ACT
		[ ]	GETACTINDEX /D3/DATA/ACT
		[ ]	GETACT 		/D3/DATA/ACT/{ACTID}

			D3 ARTISAN AND RECIPE
		[ ]	GETARTISAN 	/D3/DATA/ARTISAN/{ARTISANSLUG}
		[ ]	GETRECIPE 	/D3/DATA/ARTISAN/{ARTISANSLUG}/RECIPE/{RECIPESLUG}

			D3 FOLLOWER
		[ ]	GETFOLLOWER /D3/DATA/FOLLOWER/{FOLLOWERSLUG}

			D3 CHARACTER CLASS AND SKILL
		[ ]	GETCHARACTERCLASS 	/D3/DATA/HERO/{CLASSSLUG}
		[ ]	GETAPISKILL 		/D3/DATA/HERO/{CLASSSLUG}/SKILL/{SKILLSLUG}

			D3 ITEM TYPE
		[ ]	GETITEMTYPEINDEX 	/D3/DATA/ITEM-TYPE
		[ ]	GETITEMTYPE 		/D3/DATA/ITEM-TYPE/{ITEMTYPESLUG}

			D3 ITEM
		[ ]	GETITEM /D3/DATA/ITEM/{ITEMSLUGANDID}

			D3 PROFILE
		[ ]	GETAPIACCOUNT 				/D3/PROFILE/{ACCOUNT}/
		[ ]	GETAPIHERO 					/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}
		[ ]	GETAPIDETAILEDHEROITEMS 	/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/ITEMS
		[ ]	GETAPIDETAILEDFOLLOWERITEMS /D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/FOLLOWER-ITEMS

	*/

	/*
			D3 GAME DATA API
			--------------------------------------------------------------------
			D3
		[ ]	SEASON INDEX 		/DATA/D3/SEASON/
		[ ]	SEASON 				/DATA/D3/SEASON/:ID
		[ ]	SEASON LEADERBOARD 	/DATA/D3/SEASON/:ID/LEADERBOARD/:LEADERBOARD
		[ ]	ERA INDEX 			/DATA/D3/ERA/
		[ ]	ERA 				/DATA/D3/ERA/:ID
		[ ]	ERA LEADERBOARD 	/DATA/D3/ERA/:ID/LEADERBOARD/:LEADERBOARD

	*/
)

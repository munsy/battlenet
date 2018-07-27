package d3

import (
	"strconv"

	"github.com/munsy/battlenet/regions"
)

var (
	//  API
	endpointD3   = func(r regions.Region) string { return r.API() + "d3/" }
	endpointData = func(r regions.Region) string { return endpointD3(r) + "data/" }

	//  ACT
	endpointActs = func(r regions.Region) string { return endpointData(r) + "act" }                                // GETACTINDEX /D3/DATA/ACT
	endpointAct  = func(r regions.Region, actID int) string { return endpointActs(r) + "/" + strconv.Itoa(actID) } // GETACT 	 /D3/DATA/ACT/{ACTID}

	//  ARTISAN AND RECIPE
	artisan         = func(r regions.Region) string { return endpointData(r) + "artisan/" }
	endpointArtisan = func(r regions.Region, artisanSlug string) string { return artisan(r) + artisanSlug } // GETARTISAN 	/D3/DATA/ARTISAN/{ARTISANSLUG}
	endpointRecipe  = func(r regions.Region, artisanSlug, recipeSlug string) string {                       // GETRECIPE 	/D3/DATA/ARTISAN/{ARTISANSLUG}/RECIPE/{RECIPESLUG}
		return endpointArtisan(r, artisanSlug) + "/" + recipeSlug
	}

	//  FOLLOWER
	follower         = func(r regions.Region) string { return endpointData(r) + "follower/" }
	endpointFollower = func(r regions.Region, followerSlug string) string { return follower(r) + followerSlug } // 	GETFOLLOWER /D3/DATA/FOLLOWER/{FOLLOWERSLUG}

	//  CHARACTER CLASS AND SKILL
	characterClass         = func(r regions.Region) string { return endpointData(r) + "hero/" }
	endpointCharacterClass = func(r regions.Region, classSlug string) string { return characterClass(r) + classSlug } // GETCHARACTERCLASS 	/D3/DATA/HERO/{CLASSSLUG}
	endpointSkill          = func(r regions.Region, classSlug, skillSlug string) string {                             // GETAPISKILL 		/D3/DATA/HERO/{CLASSSLUG}/SKILL/{SKILLSLUG}
		return endpointCharacterClass(r, classSlug) + "/skill/" + skillSlug
	}

	//  ITEM TYPE
	endpointItemTypeIndex = func(r regions.Region) string { return endpointData(r) + "item-type" } // GETITEMTYPEINDEX 	/D3/DATA/ITEM-TYPE
	endpointItemType      = func(r regions.Region, itemTypeSlug string) string {
		return endpointItemTypeIndex(r) + "/" + itemTypeSlug
	} // GETITEMTYPE 		/D3/DATA/ITEM-TYPE/{ITEMTYPESLUG}

	//  ITEM
	item         = func(r regions.Region) string { return endpointData(r) + "item/" }
	endpointItem = func(r regions.Region, itemSlugAndID string) string { return item(r) + itemSlugAndID } // GETITEM /D3/DATA/ITEM/{ITEMSLUGANDID}

	//  PROFILE
	profile         = func(r regions.Region) string { return endpointD3(r) + "profile/" }
	endpointAccount = func(r regions.Region, account string) string { return profile(r) + account } // 	GETAPIACCOUNT	/D3/PROFILE/{ACCOUNT}/
	endpointHero    = func(r regions.Region, account string, heroID int) string {                   // 	GETAPIHERO 		/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}
		return endpointAccount(r, account) + "/hero/" + strconv.Itoa(heroID)
	}
	endpointDetailedHeroItems = func(r regions.Region, account string, heroID int) string { // GETAPIDETAILEDHEROITEMS 	/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/ITEMS
		return endpointHero(r, account, heroID) + "/items"
	}
	endpointDetailedFollowerItems = func(r regions.Region, account string, heroID int) string { // GETAPIDETAILEDFOLLOWERITEMS /D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/FOLLOWER-ITEMS
		return endpointHero(r, account, heroID) + "/follower-items"
	}

	//  GAME DATA API
	d3Data                    = func(r regions.Region) string { return r.API() + "/data/" }
	dataD3                    = func(r regions.Region) string { return d3Data(r) + "d3/" }
	endpointSeasonIndex       = func(r regions.Region) string { return dataD3(r) + "season/" }                             // SEASON INDEX 			/DATA/D3/SEASON/
	endpointSeason            = func(r regions.Region, id int) string { return endpointSeasonIndex(r) + strconv.Itoa(id) } // SEASON 				/DATA/D3/SEASON/:ID
	endpointSeasonLeaderboard = func(r regions.Region, id int, leaderboard string) string {                                // SEASON LEADERBOARD 	/DATA/D3/SEASON/:ID/LEADERBOARD/:LEADERBOARD
		return endpointSeason(r, id) + "/leaderboard/" + leaderboard
	}
	endpointEraIndex       = func(r regions.Region) string { return dataD3(r) + "era/" }                             // ERA INDEX 			/DATA/D3/ERA/
	endpointEra            = func(r regions.Region, id int) string { return endpointEraIndex(r) + strconv.Itoa(id) } // ERA 				/DATA/D3/ERA/:ID
	endpointEraLeaderboard = func(r regions.Region, id int, leaderboard string) string {                             // ERA LEADERBOARD 	/DATA/D3/ERA/:ID/LEADERBOARD/:LEADERBOARD
		return endpointEra(r, id) + "/leaderboard/" + leaderboard
	}
)

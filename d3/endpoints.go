package d3

import (
	"strconv"

	"github.com/munsy/gobattlenet/regions"
)

var (
	//  API
	endpointD3   = func(r battlenet.Region) string { return r.API() + "d3/" }
	endpointData = func(r battlenet.Region) string { return endpointD3(r) + "data/" }

	//  ACT
	acts = func(r battlenet.Region) string { return endpointData(r) + "act" }                        // GETACTINDEX /D3/DATA/ACT
	act  = func(r battlenet.Region, actID int) string { return acts(r) + "/" + strconv.Itoa(actID) } // GETACT 	 /D3/DATA/ACT/{ACTID}

	//  ARTISAN AND RECIPE
	endpointArtisan = func(r battlenet.Region) string { return endpointData(r) + "artisan/" }
	artisan         = func(r battlenet.Region, artisanSlug string) string { return endpointArtisan(r) + artisanSlug } // GETARTISAN 	/D3/DATA/ARTISAN/{ARTISANSLUG}
	recipe          = func(r battlenet.Region, artisanSlug, recipeSlug string) string {                               // GETRECIPE 	/D3/DATA/ARTISAN/{ARTISANSLUG}/RECIPE/{RECIPESLUG}
		return Artisan(r, artisanSlug) + "/" + recipeSlug
	}

	//  FOLLOWER
	endpointFollower = func(r battlenet.Region) string { return endpointData(r) + "follower/" }
	follower         = func(r battlenet.Region, followerSlug string) string { return endpointFollower(r) + followerSlug } // 	GETFOLLOWER /D3/DATA/FOLLOWER/{FOLLOWERSLUG}

	//  CHARACTER CLASS AND SKILL
	endpointCharacterClass = func(r battlenet.Region) string { return endpointData(r) + "hero/" }
	characterClass         = func(r battlenet.Region, classSlug string) string { return endpointCharacterClass(r) + classSlug } // GETCHARACTERCLASS 	/D3/DATA/HERO/{CLASSSLUG}
	skill                  = func(r battlenet.Region, classSlug, skillSlug string) string {                                     // GETAPISKILL 		/D3/DATA/HERO/{CLASSSLUG}/SKILL/{SKILLSLUG}
		return characterClass(r, classSlug) + "/skill/" + skillSlug
	}

	//  ITEM TYPE
	itemTypeIndex = func(r battlenet.Region) string { return endpointData(r) + "item-type" }                              // GETITEMTYPEINDEX 	/D3/DATA/ITEM-TYPE
	itemType      = func(r battlenet.Region, itemTypeSlug string) string { return itemTypeIndex(r) + "/" + itemTypeSlug } // GETITEMTYPE 		/D3/DATA/ITEM-TYPE/{ITEMTYPESLUG}

	//  ITEM
	endpointItem = func(r battlenet.Region) string { return endpointData(r) + "item/" }
	item         = func(r battlenet.Region, itemSlugAndID string) string { return endpointItem(r) + itemSlugAndID } // GETITEM /D3/DATA/ITEM/{ITEMSLUGANDID}

	//  PROFILE
	endpointProfile = endpointD3 + "profile/"
	account         = func(r battlenet.Region, account string) string { return endpointProfile(r) + account } // 	GETAPIACCOUNT	/D3/PROFILE/{ACCOUNT}/
	hero            = func(r battlenet.Region, acct string, heroID int) string {                              // 	GETAPIHERO 		/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}
		return account(r, acct) + "/hero/" + strconv.Itoa(heroID)
	}
	detailedHeroItems = func(r battlenet.Region, account string, heroID int) string { // GETAPIDETAILEDHEROITEMS 	/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/ITEMS
		return hero(r, account, heroID) + "/items"
	}
	detailedFollowerItems = func(r battlenet.Region, account string, heroID int) string { // GETAPIDETAILEDFOLLOWERITEMS /D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/FOLLOWER-ITEMS
		return hero(r, account, heroID) + "/follower-items"
	}

	//  GAME DATA API
	endpointD3Data    = func(r battlenet.Region) string { return r.API() + "/data/" }
	endpointDataD3    = func(r battlenet.Region) string { return endpointD3Data(r) + "d3/" }
	seasonIndex       = func(r battlenet.Region) string { return endpointDataD3(r) + "season/" }             // SEASON INDEX 		/DATA/D3/SEASON/
	season            = func(r battlenet.Region, id int) string { return seasonIndex(r) + strconv.Itoa(id) } // SEASON 				/DATA/D3/SEASON/:ID
	seasonLeaderboard = func(r battlenet.Region, id int, leaderboard string) string {                        // SEASON LEADERBOARD 	/DATA/D3/SEASON/:ID/LEADERBOARD/:LEADERBOARD
		return season(r, id) + "/leaderboard/" + leaderboard
	}
	eraIndex       = func(r battlenet.Region) string { return endpointDataD3(r) + "era/" }             // ERA INDEX 		/DATA/D3/ERA/
	era            = func(r battlenet.Region, id int) string { return eraIndex(r) + strconv.Itoa(id) } // ERA 				/DATA/D3/ERA/:ID
	eraLeaderboard = func(r battlenet.Region, id int, leaderboard string) string {                     // ERA LEADERBOARD 	/DATA/D3/ERA/:ID/LEADERBOARD/:LEADERBOARD
		return era(r, id) + "/leaderboard/" + leaderboard
	}
)

package d3

import (
	"strconv"

	"github.com/munsy/gobattlenet/regions"
)

var (
	//  API
	endpointD3   = regions.API() + "d3/"
	endpointData = endpointD3 + "data/"

	//  ACT
	// 	GETACTINDEX /D3/DATA/ACT
	Acts = endpointData + "act"
	// 	GETACT 		/D3/DATA/ACT/{ACTID}
	Act = func(actID int) string { return Acts + "/" + strconv.Itoa(actID) }

	//  ARTISAN AND RECIPE
	endpointArtisan = endpointData + "artisan/"
	// 	GETARTISAN 	/D3/DATA/ARTISAN/{ARTISANSLUG}
	Artisan = func(artisanSlug string) string { return endpointArtisan + artisanSlug }
	// 	GETRECIPE 	/D3/DATA/ARTISAN/{ARTISANSLUG}/RECIPE/{RECIPESLUG}
	Recipe = func(artisanSlug, recipeSlug string) string { return Artisan(artisanSlug) + "/" + recipeSlug }

	//  FOLLOWER
	endpointFollower = endpointData + "follower/"
	// 	GETFOLLOWER /D3/DATA/FOLLOWER/{FOLLOWERSLUG}
	Follower = func(followerSlug string) string { return endpointFollower + followerSlug }

	//  CHARACTER CLASS AND SKILL
	endpointCharacterClass = endpointData + "hero/"
	// 	GETCHARACTERCLASS 	/D3/DATA/HERO/{CLASSSLUG}
	CharacterClass = func(classSlug string) string { return endpointCharacterClass + classSlug }
	// 	GETAPISKILL 		/D3/DATA/HERO/{CLASSSLUG}/SKILL/{SKILLSLUG}
	Skill = func(classSlug, skillSlug string) string {
		return CharacterClass(classSlug) + "/skill/" + skillSlug
	}

	//  ITEM TYPE
	// 	GETITEMTYPEINDEX 	/D3/DATA/ITEM-TYPE
	ItemTypeIndex = endpointData + "item-type"
	// 	GETITEMTYPE 		/D3/DATA/ITEM-TYPE/{ITEMTYPESLUG}
	ItemType = func(itemTypeSlug string) string { return ItemTypeIndex + "/" + itemTypeSlug }

	//  ITEM
	endpointItem = endpointData + "item/"
	// 	GETITEM /D3/DATA/ITEM/{ITEMSLUGANDID}
	Item = func(itemSlugAndID string) string { return endpointItem + itemSlugAndID }

	//  PROFILE
	endpointProfile = endpointD3 + "profile/"
	// 	GETAPIACCOUNT 				/D3/PROFILE/{ACCOUNT}/
	Account = func(account string) string { return endpointProfile + account }
	// 	GETAPIHERO 					/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}
	Hero = func(account string, heroID int) string {
		return Account(account) + "/hero/" + strconv.Itoa(heroID)
	}
	// 	GETAPIDETAILEDHEROITEMS 	/D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/ITEMS
	DetailedHeroItems = func(account string, heroID int) string { return Hero(account, heroID) + "/items" }
	// 	GETAPIDETAILEDFOLLOWERITEMS /D3/PROFILE/{ACCOUNT}/HERO/{HEROID}/FOLLOWER-ITEMS
	DetailedFollowerItems = func(account string, heroID int) string { return Hero(account, heroID) + "/follower-items" }

	//  GAME DATA API
	endpointD3Data = "/data/"
	endpointDataD3 = endpointD3Data + "d3/"
	// 	SEASON INDEX 		/DATA/D3/SEASON/
	SeasonIndex = endpointDataD3 + "season/"
	// 	SEASON 				/DATA/D3/SEASON/:ID
	Season = func(id int) string { return SeasonIndex + strconv.Itoa(id) }
	// 	SEASON LEADERBOARD 	/DATA/D3/SEASON/:ID/LEADERBOARD/:LEADERBOARD
	SeasonLeaderboard = func(id int, leaderboard string) string { return Season(id) + "/leaderboard/" + leaderboard }
	// 	ERA INDEX 			/DATA/D3/D3ERA/
	EraIndex = endpointDataD3 + "era/"
	// 	ERA 				/DATA/D3/ERA/:ID
	Era = func(id int) string { return EraIndex + strconv.Itoa(id) }
	// 	ERA LEADERBOARD 	/DATA/D3/ERA/:ID/LEADERBOARD/:LEADERBOARD
	EraLeaderboard = func(id int, leaderboard string) string { return Era(id) + "/leaderboard/" + leaderboard }
)

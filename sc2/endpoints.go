package sc2

import (
	"strconv"

	"github.com/munsy/gobattlenet/regions"
)

var (
	// SC2 API
	sc2 = regions.API() + "sc2/"
	// PROFILE
	endpointProfile = sc2 + "profile/"
	user            = endpointProfile + "user"                                              // SC2 OAUTH PROFILE 	/SC2/PROFILE/USER
	profile         = func(r battlenet.Region, profileID int, region, name string) string { // PROFILE				/SC2/PROFILE/:ID/:REGION/:NAME/
		return endpointProfile + strconv.Itoa(profileID) + "/" + region + "/" + name + "/"
	}
	ladderProfile = func(profileID int, region, name string) string { // LADDERS 	/SC2/PROFILE/:ID/:REGION/:NAME/LADDERS
		return profile(profileID, region, name) + "ladders"
	}

	matchHistory = func(profileID int, region, name string) string { // MATCH HISTORY 	/SC2/PROFILE/:ID/:REGION/:NAME/MATCHES
		return profile(profileID, region, name) + "matches"
	}

	// LADDER
	endpointLadder = sc2 + "ladder/"
	ladder         = func(id int) string { return endpointLadder + strconv.Itoa(id) } // LADDER /SC2/LADDER/:ID

	// DATA RESOURCES
	endpointData = sc2 + "data/"
	achievements = endpointData + "achievements" // ACHIEVEMENTS 	/SC2/DATA/ACHIEVEMENTS
	rewards      = endpointData + "rewards"      // REWARDS 		/SC2/DATA/REWARDS
)

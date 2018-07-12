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
	User            = endpointProfile + "user"                          // SC2 OAUTH PROFILE /SC2/PROFILE/USER
	Profile         = func(profileID int, region, name string) string { // PROFILE	/SC2/PROFILE/:ID/:REGION/:NAME/
		return endpointProfile + strconv.Itoa(profileID) + "/" + region + "/" + name + "/"
	}
	LadderProfile = func(profileID int, region, name string) string { // LADDERS 	/SC2/PROFILE/:ID/:REGION/:NAME/LADDERS
		return Profile(profileID, region, name) + "ladders"
	}

	MatchHistory = func(profileID int, region, name string) string { // MATCH HISTORY 	/SC2/PROFILE/:ID/:REGION/:NAME/MATCHES
		return Profile(profileID, region, name) + "matches"
	}

	// LADDER
	endpointLadder = sc2 + "ladder/"
	Ladder         = func(id int) string { return endpointLadder + strconv.Itoa(id) } // LADDER /SC2/LADDER/:ID

	// DATA RESOURCES
	endpointData = sc2 + "data/"
	Achievements = endpointData + "achievements" // ACHIEVEMENTS 	/SC2/DATA/ACHIEVEMENTS
	Rewards      = endpointData + "rewards"      // REWARDS 			/SC2/DATA/REWARDS
)

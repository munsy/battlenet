package sc2

import (
	"strconv"

	"github.com/munsy/gobattlenet/regions"
)

var (
	// SC2 API
	endpointSc2 = func(r regions.Region) string { return r.API() + "sc2/" }
	// PROFILE
	profile         = func(r regions.Region) string { return endpointSc2(r) + "profile/" }
	user            = func(r regions.Region) string { return profile(r) + "user" }        // SC2 OAUTH PROFILE 	/SC2/PROFILE/USER
	endpointProfile = func(r regions.Region, profileID int, region, name string) string { // PROFILE				/SC2/PROFILE/:ID/:REGION/:NAME/
		return profile(r) + strconv.Itoa(profileID) + "/" + region + "/" + name + "/"
	}
	endpointLadderProfile = func(r regions.Region, profileID int, region, name string) string { // LADDERS 	/SC2/PROFILE/:ID/:REGION/:NAME/LADDERS
		return endpointProfile(r, profileID, region, name) + "ladders"
	}

	endpointMatchHistory = func(r regions.Region, profileID int, region, name string) string { // MATCH HISTORY 	/SC2/PROFILE/:ID/:REGION/:NAME/MATCHES
		return endpointProfile(r, profileID, region, name) + "matches"
	}

	// LADDER
	ladder         = func(r regions.Region) string { return endpointSc2(r) + "ladder/" }
	endpointLadder = func(r regions.Region, id int) string { return ladder(r) + strconv.Itoa(id) } // LADDER /SC2/LADDER/:ID

	// DATA RESOURCES
	endpointData         = func(r regions.Region) string { return endpointSc2(r) + "data/" }
	endpointAchievements = func(r regions.Region) string { return endpointData(r) + "achievements" } // ACHIEVEMENTS 	/SC2/DATA/ACHIEVEMENTS
	endpointRewards      = func(r regions.Region) string { return endpointData(r) + "rewards" }      // REWARDS 		/SC2/DATA/REWARDS
)

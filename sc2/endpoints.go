package sc2

import (
	"strconv"

	"github.com/munsy/gobattlenet"
)

var (
	// SC2 API
	sc2 = func(r battlenet.Region) string { return r.API() + "sc2/" }
	// PROFILE
	endpointProfile = func(r battlenet.Region) string { return sc2(r) + "profile/" }
	user            = func(r battlenet.Region) string { return endpointProfile(r) + "user" } // SC2 OAUTH PROFILE 	/SC2/PROFILE/USER
	profile         = func(r battlenet.Region, profileID int, region, name string) string {  // PROFILE				/SC2/PROFILE/:ID/:REGION/:NAME/
		return endpointProfile(r) + strconv.Itoa(profileID) + "/" + region + "/" + name + "/"
	}
	ladderProfile = func(r battlenet.Region, profileID int, region, name string) string { // LADDERS 	/SC2/PROFILE/:ID/:REGION/:NAME/LADDERS
		return profile(r, profileID, region, name) + "ladders"
	}

	matchHistory = func(r battlenet.Region, profileID int, region, name string) string { // MATCH HISTORY 	/SC2/PROFILE/:ID/:REGION/:NAME/MATCHES
		return profile(r, profileID, region, name) + "matches"
	}

	// LADDER
	endpointLadder = func(r battlenet.Region) string { return sc2(r) + "ladder/" }
	ladder         = func(r battlenet.Region, id int) string { return endpointLadder(r) + strconv.Itoa(id) } // LADDER /SC2/LADDER/:ID

	// DATA RESOURCES
	endpointData = func(r battlenet.Region) string { return sc2(r) + "data/" }
	achievements = func(r battlenet.Region) string { return endpointData(r) + "achievements" } // ACHIEVEMENTS 	/SC2/DATA/ACHIEVEMENTS
	rewards      = func(r battlenet.Region) string { return endpointData(r) + "rewards" }      // REWARDS 		/SC2/DATA/REWARDS
)

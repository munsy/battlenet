package endpoints

import (
	"strconv"
)

var (
	// SC2 API
	sc2 = API() + "sc2/"
	// PROFILE
	endpointSc2Profile = sc2 + "profile/"
	Sc2User            = endpointSc2Profile + "user"                       // SC2 OAUTH PROFILE /SC2/PROFILE/USER
	Sc2Profile         = func(profileID int, region, name string) string { // PROFILE	/SC2/PROFILE/:ID/:REGION/:NAME/
		return endpointSc2Profile + strconv.Itoa(profileID) + "/" + region + "/" + name + "/"
	}
	Sc2LadderProfile = func(profileID int, region, name string) string { // LADDERS 	/SC2/PROFILE/:ID/:REGION/:NAME/LADDERS
		return Sc2Profile(profileID, region, name) + "ladders"
	}

	Sc2MatchHistory = func(profileID int, region, name string) string { // MATCH HISTORY 	/SC2/PROFILE/:ID/:REGION/:NAME/MATCHES
		return Sc2Profile(profileID, region, name) + "matches"
	}

	// LADDER
	endpointSc2Ladder = sc2 + "ladder/"
	Sc2Ladder         = func(id int) string { return endpointSc2Ladder + strconv.Itoa(id) } // LADDER /SC2/LADDER/:ID

	// DATA RESOURCES
	endpointSc2Data = sc2 + "data/"
	Sc2Achievements = endpointSc2Data + "achievements" // ACHIEVEMENTS 	/SC2/DATA/ACHIEVEMENTS
	Sc2Rewards      = endpointSc2Data + "rewards"      // REWARDS 			/SC2/DATA/REWARDS
)

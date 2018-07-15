package account

import (
	"github.com/munsy/gobattlenet/regions"
)

var (
	// SC2 API
	sc2 = regions.API() + "sc2/"
	// SC2 PROFILE
	endpointSc2Profile = sc2 + "profile/"
	endpointSc2User    = endpointSc2Profile + "user" // SC2 OAUTH PROFILE 	/SC2/PROFILE/USER

	// WOW API
	wow = regions.API() + "wow/"
	//	WOW PROFILE
	endpointWowUser       = wow + "user/"
	endpointWowCharacters = User + "characters" // WOW OAUTH PROFILE /WOW/USER/CHARACTERS
)

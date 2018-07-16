package account

import (
	"github.com/munsy/gobattlenet/regions"
)

var (
	// Community Oauth Profile API
	endpointAccount = func(r battlenet.Region) string { return r.API() + "account/" }
	endpointUser    = func(r battlenet.Region) string { return endpointAccount(r) + "user" } // /ACCOUNT/USER

	// SC2 API
	sc2 = func(r battlenet.Region) string { return r.API() + "sc2/" }
	// SC2 PROFILE
	endpointSc2Profile = func(r battlenet.Region) string { return sc2(r) + "profile/" }
	endpointSc2User    = func(r battlenet.Region) string { return endpointSc2Profile(r) + "user" } // SC2 OAUTH PROFILE 	/SC2/PROFILE/USER

	// WOW API
	wow = func(r battlenet.Region) string { return r.API() + "wow/" }
	//	WOW PROFILE
	endpointWowUser       = func(r battlenet.Region) string { return wow(r) + "user/" }
	endpointWowCharacters = func(r battlenet.Region) string { return endpointWowUser(r) + "characters" } // WOW OAUTH PROFILE /WOW/USER/CHARACTERS
)

package account

import (
	"github.com/munsy/gobattlenet/pkg/regions"
)

var (
	// Community Oauth Profile API
	endpointAccount = func(r regions.Region) string { return r.API() + "account/" }
	endpointUser    = func(r regions.Region) string { return endpointAccount(r) + "user" } // /ACCOUNT/USER

	// SC2 API
	endpointSc2 = func(r regions.Region) string { return r.API() + "sc2/" }
	// SC2 PROFILE
	endpointSc2Profile = func(r regions.Region) string { return endpointSc2(r) + "profile/" }
	endpointSc2User    = func(r regions.Region) string { return endpointSc2Profile(r) + "user" } // SC2 OAUTH PROFILE 	/SC2/PROFILE/USER

	// WOW API
	endpointWow = func(r regions.Region) string { return r.API() + "wow/" }
	//	WOW PROFILE
	endpointWowUser       = func(r regions.Region) string { return endpointWow(r) + "user/" }
	endpointWowCharacters = func(r regions.Region) string { return endpointWowUser(r) + "characters" } // WOW OAUTH PROFILE /WOW/USER/CHARACTERS
)

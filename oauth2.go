// Package battlenet provides clients for accessing Blizzard Entertainment's various
// Battle.net API endpoints, including Diablo III, Starcraft II, and World of Warcraft.
//
// Official Battle.net API documentation can be found at: https://dev.battle.net/io-docs
package battlenet

import (
	"golang.org/x/oauth2"

	"github.com/munsy/gobattlenet/pkg/regions"
)

// Endpoint is Battle.net's OAuth 2.0 endpoint.
// Defaults to US if Region is not otherwise defined.
//
// For more information see:
// https://dev.battle.net/docs/read/oauth
var Endpoint = func(r regions.Region) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  r.AuthURL(),
		TokenURL: r.TokenURL(),
	}
}

// Package battlenet provides clients for communicating with Blizzard's Battle.net APIs.
package battlenet

import (
	"golang.org/x/oauth2"

	"github.com/munsy/gobattlenet/regions"
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

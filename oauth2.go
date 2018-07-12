package battlenet // import "github.com/munsy/gobattlenet"

import (
	"golang.org/x/oauth2"

	"github.com/munsy/gobattlenet/regions"
)

// Endpoint is Battle.net's OAuth 2.0 endpoint.
// Defaults to US if Region is not otherwise defined.
//
// For more information see:
// https://dev.battle.net/docs/read/oauth
var Endpoint = oauth2.Endpoint{
	AuthURL:  regions.AuthURL,
	TokenURL: regions.TokenURL,
}

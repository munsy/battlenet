// Package battlenet provides clients for accessing Blizzard Entertainment's various Battle.net
// API endpoints, including ones for Diablo III, Starcraft II, and World of Warcraft.
//
// Official Battle.net API documentation can be found at: https://dev.battle.net/io-docs
package battlenet

import (
	"github.com/munsy/gobattlenet/http"
	"github.com/munsy/gobattlenet/settings"
)

// New returns a new  client.Client for accessing the Battle.net API.
func New(s *settings.BNetSettings) (c *http.Client, err error) {
	return http.New(s)
}

// Package battlenet provides clients for accessing Blizzard Entertainment's various Battle.net
// API endpoints, including ones for Diablo III, Starcraft II, and World of Warcraft.
//
// Official Battle.net API documentation can be found at: https://dev.battle.net/io-docs
package battlenet

import (
	"github.com/munsy/gobattlenet/client/account"
	"github.com/munsy/gobattlenet/client/d3"
	"github.com/munsy/gobattlenet/client/sc2"
	"github.com/munsy/gobattlenet/client/wow"
	"github.com/munsy/gobattlenet/settings"
)

// NewAccountClient returns a new Account Client for accessing the Battle.net Account API.
func NewAccountClient(s *settings.BNetSettings) (c *account.Client, err error) {
	return account.New(s)
}

// NewD3Client returns a new D3 Client for accessing the Diablo III Battle.net API.
func NewD3Client(s *settings.BNetSettings) (c *d3.Client, err error) {
	return d3.New(s)
}

// NewSC2Client returns a new SC2 Client for accessing the Starcraft II Battle.net API.
func NewSC2Client(s *settings.BNetSettings) (c *sc2.Client, err error) {
	return sc2.New(s)
}

// NewWoWClient returns a new WoW Client for accessing the World of Warcraft Battle.net API.
func NewWoWClient(s *settings.BNetSettings) (c *wow.Client, err error) {
	return wow.New(s)
}

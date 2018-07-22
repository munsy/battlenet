package battlenet

import (
	"github.com/munsy/gobattlenet/account"
	"github.com/munsy/gobattlenet/d3"
	"github.com/munsy/gobattlenet/sc2"
	"github.com/munsy/gobattlenet/wow"
)

// NewAccountClient returns a new AccountClient for accessing the Battle.net Account API.
func NewAccountClient(args ...interface{}) (c *account.AccountClient, err error) {
	return account.New(args)
}

// NewD3Client returns a new D3Client for accessing the Diablo III API.
func NewD3Client(args ...interface{}) (c *d3.D3Client, err error) {
	return d3.New(args)
}

// NewSC2Client returns a new SC2Client for accessing the Starcraft II API.
func NewSC2Client(args ...interface{}) (c *sc2.SC2Client, err error) {
	return sc2.New(args)
}

// NewWoWClient returns a new WoWClient for accessing the World of Warcraft API.
func NewWoWClient(args ...interface{}) (c *wow.WoWClient, err error) {
	return wow.New(args)
}

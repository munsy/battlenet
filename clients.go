package battlenet

import (
	"github.com/munsy/gobattlenet/client/account"
	"github.com/munsy/gobattlenet/client/d3"
	"github.com/munsy/gobattlenet/client/sc2"
	"github.com/munsy/gobattlenet/client/wow"
)

// NewAccountClient returns a new AccountClient for accessing the Battle.net Account API.
func NewAccountClient(args ...interface{}) (c *account.Client, err error) {
	return account.New(args)
}

// NewD3Client returns a new D3Client for accessing the Diablo III API.
func NewD3Client(args ...interface{}) (c *d3.Client, err error) {
	return d3.New(args)
}

// NewSC2Client returns a new SC2Client for accessing the Starcraft II API.
func NewSC2Client(args ...interface{}) (c *sc2.Client, err error) {
	return sc2.New(args)
}

// NewWoWClient returns a new WoWClient for accessing the World of Warcraft API.
func NewWoWClient(args ...interface{}) (c *wow.Client, err error) {
	return wow.New(args)
}

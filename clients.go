package battlenet

import (
	"github.com/munsy/battlenet/http/account"
	"github.com/munsy/battlenet/http/d3"
	"github.com/munsy/battlenet/http/sc2"
	"github.com/munsy/battlenet/http/wow"
)

const version = "beta"

// AccountClient returns a new client for accessing the Battle.net Account API.
func AccountClient(s *Settings, token string) (*account.Client, error) {
	return account.New(s.Client, s.Region, s.Locale, token, version)
}

// D3Client returns a new client for accessing the Diablo III API.
func D3Client(s *Settings, key string) (*d3.Client, error) {
	return d3.New(s.Client, s.Region, s.Locale, key, version)
}

// SC2Client returns a new client for accessing the Starcraft II API.
func SC2Client(s *Settings, key string) (*sc2.Client, error) {
	return sc2.New(s.Client, s.Region, s.Locale, key, version)
}

// WoWClient returns a new client for accessing the World of Warcraft API.
func WoWClient(s *Settings, key string) (*wow.Client, error) {
	return wow.New(s.Client, s.Region, s.Locale, key, version)
}

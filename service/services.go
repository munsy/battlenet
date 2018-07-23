package service

import (
	"github.com/munsy/gobattlenet/service/account"
	"github.com/munsy/gobattlenet/service/d3"
	"github.com/munsy/gobattlenet/service/sc2"
	"github.com/munsy/gobattlenet/service/wow"
)

// Account returns a new Account service.
func Account(token string, client *http.Client) *account.Service {
	return account.New(token, client)
}

// D3 returns a new D3 service.
func D3(key string, client *http.Client) *d3.Service {
	return d3.New(key, client)
}

// Sc2 returns a new Sc2 service.
func Sc2(key string, client *http.Client) *sc2.Service {
	return sc2.New(key, client)
}

// WoW returns a new WoW service.
func WoW(key string, client *http.Client) *wow.Serivce {
	return wow.New(key, client)
}

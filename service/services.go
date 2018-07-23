package service

import (
	"github.com/munsy/gobattlenet/service/account"
	"github.com/munsy/gobattlenet/service/d3"
	"github.com/munsy/gobattlenet/service/sc2"
	"github.com/munsy/gobattlenet/service/wow"
)

// Account returns a new Account service.
func Account(token string) *account.Service {
	return account.Service(token)
}

// D3 returns a new D3 service.
func D3(key string) *d3.Service {
	return d3.Service(key)
}

// Sc2 returns a new Sc2 service.
func Sc2(key string) *sc2.Service {
	return sc2.Service(key)
}

// WoW returns a new WoW service.
func WoW(key string) *wow.Serivce {
	return wow.Service(key)
}

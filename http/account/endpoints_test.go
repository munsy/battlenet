package account

import (
	"testing"

	"github.com/munsy/battlenet/regions"
)

func TestEndpointsUS(t *testing.T) {
	r := regions.US
	if endpointAccount(r) != "https://us.api.battle.net/account/" {
		t.Fatal("wrong endpoint: " + endpointAccount(r))
	}
	if endpointUser(r) != "https://us.api.battle.net/account/user" {
		t.Fatal("wrong endpoint: " + endpointUser(r))
	}
	if endpointSc2(r) != "https://us.api.battle.net/sc2/" {
		t.Fatal("wrong endpoint: " + endpointSc2(r))
	}
	if endpointSc2Profile(r) != "https://us.api.battle.net/sc2/profile/" {
		t.Fatal("wrong endpoint: " + endpointSc2Profile(r))
	}
	if endpointSc2User(r) != "https://us.api.battle.net/sc2/profile/user" {
		t.Fatal("wrong endpoint: " + endpointSc2User(r))
	}
	if endpointWow(r) != "https://us.api.battle.net/wow/" {
		t.Fatal("wrong endpoint: " + endpointWow(r))
	}
	if endpointWowUser(r) != "https://us.api.battle.net/wow/user/" {
		t.Fatal("wrong endpoint: " + endpointWowUser(r))
	}
	if endpointWowCharacters(r) != "https://us.api.battle.net/wow/user/characters" {
		t.Fatal("wrong endpoint: " + endpointWowCharacters(r))
	}
}

func TestEndpointsEU(t *testing.T) {
	r := regions.EU
	if endpointAccount(r) != "https://eu.api.battle.net/account/" {
		t.Fatal("wrong endpoint: " + endpointAccount(r))
	}
	if endpointUser(r) != "https://eu.api.battle.net/account/user" {
		t.Fatal("wrong endpoint: " + endpointUser(r))
	}
	if endpointSc2(r) != "https://eu.api.battle.net/sc2/" {
		t.Fatal("wrong endpoint: " + endpointSc2(r))
	}
	if endpointSc2Profile(r) != "https://eu.api.battle.net/sc2/profile/" {
		t.Fatal("wrong endpoint: " + endpointSc2Profile(r))
	}
	if endpointSc2User(r) != "https://eu.api.battle.net/sc2/profile/user" {
		t.Fatal("wrong endpoint: " + endpointSc2User(r))
	}
	if endpointWow(r) != "https://eu.api.battle.net/wow/" {
		t.Fatal("wrong endpoint: " + endpointWow(r))
	}
	if endpointWowUser(r) != "https://eu.api.battle.net/wow/user/" {
		t.Fatal("wrong endpoint: " + endpointWowUser(r))
	}
	if endpointWowCharacters(r) != "https://eu.api.battle.net/wow/user/characters" {
		t.Fatal("wrong endpoint: " + endpointWowCharacters(r))
	}
}

func TestEndpointsKR(t *testing.T) {
	r := regions.KR
	if endpointAccount(r) != "https://kr.api.battle.net/account/" {
		t.Fatal("wrong endpoint: " + endpointAccount(r))
	}
	if endpointUser(r) != "https://kr.api.battle.net/account/user" {
		t.Fatal("wrong endpoint: " + endpointUser(r))
	}
	if endpointSc2(r) != "https://kr.api.battle.net/sc2/" {
		t.Fatal("wrong endpoint: " + endpointSc2(r))
	}
	if endpointSc2Profile(r) != "https://kr.api.battle.net/sc2/profile/" {
		t.Fatal("wrong endpoint: " + endpointSc2Profile(r))
	}
	if endpointSc2User(r) != "https://kr.api.battle.net/sc2/profile/user" {
		t.Fatal("wrong endpoint: " + endpointSc2User(r))
	}
	if endpointWow(r) != "https://kr.api.battle.net/wow/" {
		t.Fatal("wrong endpoint: " + endpointWow(r))
	}
	if endpointWowUser(r) != "https://kr.api.battle.net/wow/user/" {
		t.Fatal("wrong endpoint: " + endpointWowUser(r))
	}
	if endpointWowCharacters(r) != "https://kr.api.battle.net/wow/user/characters" {
		t.Fatal("wrong endpoint: " + endpointWowCharacters(r))
	}
}

func TestEndpointsTW(t *testing.T) {
	r := regions.TW
	if endpointAccount(r) != "https://tw.api.battle.net/account/" {
		t.Fatal("wrong endpoint: " + endpointAccount(r))
	}
	if endpointUser(r) != "https://tw.api.battle.net/account/user" {
		t.Fatal("wrong endpoint: " + endpointUser(r))
	}
	if endpointSc2(r) != "https://tw.api.battle.net/sc2/" {
		t.Fatal("wrong endpoint: " + endpointSc2(r))
	}
	if endpointSc2Profile(r) != "https://tw.api.battle.net/sc2/profile/" {
		t.Fatal("wrong endpoint: " + endpointSc2Profile(r))
	}
	if endpointSc2User(r) != "https://tw.api.battle.net/sc2/profile/user" {
		t.Fatal("wrong endpoint: " + endpointSc2User(r))
	}
	if endpointWow(r) != "https://tw.api.battle.net/wow/" {
		t.Fatal("wrong endpoint: " + endpointWow(r))
	}
	if endpointWowUser(r) != "https://tw.api.battle.net/wow/user/" {
		t.Fatal("wrong endpoint: " + endpointWowUser(r))
	}
	if endpointWowCharacters(r) != "https://tw.api.battle.net/wow/user/characters" {
		t.Fatal("wrong endpoint: " + endpointWowCharacters(r))
	}
}

func TestEndpointsSEA(t *testing.T) {
	r := regions.SEA
	if endpointAccount(r) != "https://sea.api.battle.net/account/" {
		t.Fatal("wrong endpoint: " + endpointAccount(r))
	}
	if endpointUser(r) != "https://sea.api.battle.net/account/user" {
		t.Fatal("wrong endpoint: " + endpointUser(r))
	}
	if endpointSc2(r) != "https://sea.api.battle.net/sc2/" {
		t.Fatal("wrong endpoint: " + endpointSc2(r))
	}
	if endpointSc2Profile(r) != "https://sea.api.battle.net/sc2/profile/" {
		t.Fatal("wrong endpoint: " + endpointSc2Profile(r))
	}
	if endpointSc2User(r) != "https://sea.api.battle.net/sc2/profile/user" {
		t.Fatal("wrong endpoint: " + endpointSc2User(r))
	}
	if endpointWow(r) != "https://sea.api.battle.net/wow/" {
		t.Fatal("wrong endpoint: " + endpointWow(r))
	}
	if endpointWowUser(r) != "https://sea.api.battle.net/wow/user/" {
		t.Fatal("wrong endpoint: " + endpointWowUser(r))
	}
	if endpointWowCharacters(r) != "https://sea.api.battle.net/wow/user/characters" {
		t.Fatal("wrong endpoint: " + endpointWowCharacters(r))
	}
}

func TestEndpointsCN(t *testing.T) {
	r := regions.CN
	if endpointAccount(r) != "https://api.battlenet.com.cn/account/" {
		t.Fatal("wrong endpoint: " + endpointAccount(r))
	}
	if endpointUser(r) != "https://api.battlenet.com.cn/account/user" {
		t.Fatal("wrong endpoint: " + endpointUser(r))
	}
	if endpointSc2(r) != "https://api.battlenet.com.cn/sc2/" {
		t.Fatal("wrong endpoint: " + endpointSc2(r))
	}
	if endpointSc2Profile(r) != "https://api.battlenet.com.cn/sc2/profile/" {
		t.Fatal("wrong endpoint: " + endpointSc2Profile(r))
	}
	if endpointSc2User(r) != "https://api.battlenet.com.cn/sc2/profile/user" {
		t.Fatal("wrong endpoint: " + endpointSc2User(r))
	}
	if endpointWow(r) != "https://api.battlenet.com.cn/wow/" {
		t.Fatal("wrong endpoint: " + endpointWow(r))
	}
	if endpointWowUser(r) != "https://api.battlenet.com.cn/wow/user/" {
		t.Fatal("wrong endpoint: " + endpointWowUser(r))
	}
	if endpointWowCharacters(r) != "https://api.battlenet.com.cn/wow/user/characters" {
		t.Fatal("wrong endpoint: " + endpointWowCharacters(r))
	}
}

package regions

import "testing"

func TestAuthURL(t *testing.T) {
	if US.AuthURL() != "https://us.battle.net/oauth/authorize" {
		t.Fatal("bad url: " + US.AuthURL())
	}
	if EU.AuthURL() != "https://eu.battle.net/oauth/authorize" {
		t.Fatal("bad url: " + EU.AuthURL())
	}
	if KR.AuthURL() != "https://kr.battle.net/oauth/authorize" {
		t.Fatal("bad url: " + KR.AuthURL())
	}
	if TW.AuthURL() != "https://tw.battle.net/oauth/authorize" {
		t.Fatal("bad url: " + TW.AuthURL())
	}
	if SEA.AuthURL() != "https://sea.battle.net/oauth/authorize" {
		t.Fatal("bad url: " + SEA.AuthURL())
	}
	if CN.AuthURL() != "https://www.battlenet.com.cn/oauth/authorize" {
		t.Fatal("bad url: " + CN.AuthURL())
	}
}

func TestTokenURL(t *testing.T) {
	if US.TokenURL() != "https://us.battle.net/oauth/token" {
		t.Fatal("bad url: " + US.TokenURL())
	}
	if EU.TokenURL() != "https://eu.battle.net/oauth/token" {
		t.Fatal("bad url: " + EU.TokenURL())
	}
	if KR.TokenURL() != "https://kr.battle.net/oauth/token" {
		t.Fatal("bad url: " + KR.TokenURL())
	}
	if TW.TokenURL() != "https://tw.battle.net/oauth/token" {
		t.Fatal("bad url: " + TW.TokenURL())
	}
	if SEA.TokenURL() != "https://sea.battle.net/oauth/token" {
		t.Fatal("bad url: " + SEA.TokenURL())
	}
	if CN.TokenURL() != "https://www.battlenet.com.cn/oauth/token" {
		t.Fatal("bad url: " + CN.TokenURL())
	}
}

func TestAPI(t *testing.T) {
	if US.API() != "https://us.api.battle.net/" {
		t.Fatal("bad url: " + US.API())
	}
	if EU.API() != "https://eu.api.battle.net/" {
		t.Fatal("bad url: " + EU.API())
	}
	if KR.API() != "https://kr.api.battle.net/" {
		t.Fatal("bad url: " + KR.API())
	}
	if TW.API() != "https://tw.api.battle.net/" {
		t.Fatal("bad url: " + TW.API())
	}
	if SEA.API() != "https://sea.api.battle.net/" {
		t.Fatal("bad url: " + SEA.API())
	}
	if CN.API() != "https://api.battlenet.com.cn/" {
		t.Fatal("bad url: " + CN.API())
	}
}

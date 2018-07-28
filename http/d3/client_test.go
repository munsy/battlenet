package d3

import (
	"net/http"
	"testing"

	"github.com/munsy/battlenet/locale"
	"github.com/munsy/battlenet/regions"
)

func TestNew(t *testing.T) {
	c, err := New(http.DefaultClient, regions.US, locale.AmericanEnglish, "no-key", "test")

	if nil != err {
		t.Fatal(err.Error())
	}

	if nil == c {
		t.Fatal("nil client")
	}
}

func TestNil(t *testing.T) {
	c, err := New(nil, regions.US, locale.AmericanEnglish, "no-key", "test")

	if nil != err {
		t.Fatal(err.Error())
	}

	if nil == c.client {
		t.Fatal("nil client")
	}
}

func TestRegion(t *testing.T) {
	c, err := New(nil, regions.US, locale.AmericanEnglish, "no-key", "test")

	if nil != err {
		t.Fatal(err.Error())
	}

	if c.Locale() != "en_US" {
		t.Fatal("wrong locale")
	}
}

func TestLocale(t *testing.T) {
	c, err := New(nil, regions.US, locale.AmericanEnglish, "no-key", "test")

	if nil != err {
		t.Fatal(err.Error())
	}

	if c.UserAgent() != "GoBattleNet/test" {
		t.Fatal("wrong version")
	}
}

func TestUserAgent(t *testing.T) {
	c, err := New(nil, regions.US, locale.AmericanEnglish, "no-key", "test")

	if nil != err {
		t.Fatal(err.Error())
	}

	if c.UserAgent() != "GoBattleNet/test" {
		t.Fatal("wrong version")
	}
}

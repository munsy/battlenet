package wow

import (
	"testing"

	"github.com/munsy/battlenet/locale"
	"github.com/munsy/battlenet/regions"
)

func TestNew(t *testing.T) {
	c, err := New(nil, regions.US, locale.AmericanEnglish, "no-key", "test")

	if nil != err {
		t.Fatal(err.Error())
	}

	if nil == c.client {
		t.Fatal("nil client")
	}

	if c.UserAgent() != "GoBattleNet/test" {
		t.Fatal("wrong version")
	}
}

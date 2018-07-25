package battlenet

import (
	"net/http"
	"testing"
	"time"

	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

// make this better

func TestNewClient(t *testing.T) {
	settings := &Settings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
	}

	c, err := NewNewAccountClient(settings, "fake-token")

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

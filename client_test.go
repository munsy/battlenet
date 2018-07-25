package battlenet

import (
	"net/http"
	"testing"
	"time"

	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/regions"
)

// make this better

func TestNewClient(t *testing.T) {
	settings := &Settings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
	}

	c, err := AccountClient(settings, "fake-token")

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

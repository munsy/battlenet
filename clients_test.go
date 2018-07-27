package battlenet

import (
	"net/http"
	"testing"
	"time"
)

// make this better

func TestNewClient(t *testing.T) {
	settings := &Settings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: Locale.AmericanEnglish,
		Region: Regions.US,
	}

	c, err := AccountClient(settings, "fake-token")

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

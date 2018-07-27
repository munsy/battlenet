package battlenet

import (
	"net/http"
	"testing"
	"time"
)

// Test valid settings (with a fake token, we're not connecting anywhere)
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

// Nil test
func TestNilSettingsClient(t *testing.T) {
	c, err := AccountClient(nil, "fake-token")

	if nil != err {
		t.Fatal("error returned:\n" + err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

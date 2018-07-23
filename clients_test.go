package battlenet

import (
	"net/http"
	"testing"
	"time"

	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
	"github.com/munsy/gobattlenet/settings"
)

// make these better

func TestNewAccountClient(t *testing.T) {
	settings := &settings.BNetSettings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
		Key:    "",
	}

	c, err := NewAccountClient(settings)

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

func TestNewD3Client(t *testing.T) {
	settings := &settings.BNetSettings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
		Key:    "",
	}

	c, err := NewD3Client(settings)

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

func TestNewSC2Client(t *testing.T) {
	settings := &settings.BNetSettings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
		Key:    "",
	}

	c, err := NewSC2Client(settings)

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

func TestNewWoWClient(t *testing.T) {
	settings := &settings.BNetSettings{
		Client: &http.Client{Timeout: (10 * time.Second)},
		Locale: locale.AmericanEnglish,
		Region: regions.US,
		Key:    "",
	}

	c, err := NewWoWClient(settings)

	if nil != err {
		t.Fatal(err.Error())
	}
	if c == nil {
		t.Fatal("nil client")
	}
}

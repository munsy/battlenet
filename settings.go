package battlenet

import (
	"net/http"

	"github.com/munsy/battlenet/locale"
	"github.com/munsy/battlenet/regions"
)

// Settings defines settings for a Client.
type Settings struct {
	Client *http.Client
	Locale locale.Locale
	Region regions.Region
}

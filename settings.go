package battlenet

import (
	"net/http"

	"github.com/munsy/gobattlenet/locale"
	"github.com/munsy/gobattlenet/regions"
)

// Settings defines settings for a Client.
type Settings struct {
	Client *http.Client
	Locale locale.Locale
	Region regions.Region
}

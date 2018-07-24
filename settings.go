package battlenet

import (
	"net/http"

	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
)

// Settings defines settings for a Client.
type Settings struct {
	Client *http.Client
	Locale locale.Locale
	Region regions.Region
}

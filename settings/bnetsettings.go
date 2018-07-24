package settings

import (
	"net/http"

	"github.com/munsy/gobattlenet/pkg/locale"
	"github.com/munsy/gobattlenet/pkg/regions"
)

// BNetSettings defines settings for a BNetClient.
type BNetSettings struct {
	Client *http.Client
	Locale locale.Locale
	Region regions.Region
}

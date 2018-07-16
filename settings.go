package battlenet

import (
	"net/http"

	"github.com/munsy/gobattlenet"
	"github.com/munsy/gobattlenet/locale"
)

// BNetSettings defines settings for a BNetClient.
type BNetSettings struct {
	Client *http.Client
	Locale locale.Locale
	Region battlenet.Region
	Key    string
}

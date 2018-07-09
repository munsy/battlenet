package battlenet

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/munsy/gobattlenet/locale"
)

// ClientVersion defines the most up-to-date version of the GoBattleNet client.
const ClientVersion = "alpha"

// BNetClient defines the client for calling the Battle.net API.
type BNetClient struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
}

// New creates a new BNetClient. Passing different interface types can cause
// different behaviors. See function definiton for more details.
func New(args ...interface{}) (c *BNetClient, err error) {
	c = &BNetClient{
		userAgent: "GoBattleNet/" + ClientVersion,
		client:    &http.Client{Timeout: (20 * time.Second)},
		locale:    locale.AmericanEnglish,
	}

	if nil == args {
		return c, nil
	}

	for _, arg := range args {
		switch t := arg.(type) {
		case *http.Client:
			c.client = t
			return c, nil
		case locale.Locale:
			c.locale = t
			return c, nil
		case string:
			return c, nil
		case []string:
			return c, nil
		default:
			return nil, errors.New(fmt.Sprintf("Type %v is not supported.", t))
		}
	}
	panic(errors.New(fmt.Sprintf("Unresolved arguments...")))
}

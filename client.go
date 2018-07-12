package battlenet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
		client:    &http.Client{Timeout: (10 * time.Second)},
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
		case BNetSettings:
			c.client = t.Client
			c.locale = t.Locale
			return c, nil
		default:
			return nil, errors.New(fmt.Sprintf("Type %v is not supported.", t))
		}
	}
	return nil, errors.New("Unresolved constructor.")
}

// Locale gets the client's locale.
func (c *BNetClient) Locale() locale.Locale {
	return c.locale
}

// SetLocale sets the client's locale.
func (c *BNetClient) SetLocale(locale locale.Locale) {
	c.locale = locale
}

// Convert an HTTP response from a given URL to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given URL and will return an error if it fails to properly unmarshal.
func (c *BNetClient) Get(url string, v interface{}) error {
	if nil == v {
		return errors.New("No interface supplied.")
	}

	response, err := http.Get(url)
	if nil != err {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err
	}

	err = json.Unmarshal([]byte(body), &v)
	if nil != err {
		return err
	}

	return nil
}

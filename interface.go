package battlenet

import (
	"github.com/munsy/gobattlenet/locale"
)

// Client is the interface for each individual game type to implement.
type Client interface {
	// Locale gets the client's locale.
	Locale() locale.Locale

	// SetLocale sets the client's locale.
	SetLocale(locale locale.Locale)

	// SetKey sets the client's key.
	SetKey(token string)

	// UserAgent returns the client's user agent.
	UserAgent() string
}

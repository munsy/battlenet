package battlenet

// ClientVersion defines the most up-to-date version of the GoBattleNet client.
const ClientVersion = "1.0"

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

package battlenet

// Client defines the Battle.net API client interface.
type Client interface {
	// Region returns the client's current region as a string.
	Region() string

	// Locale returns the client's current locale as a string.
	Locale() string

	// UserAgent returns the client User-Agent header used in API requests.
	UserAgent() string
}

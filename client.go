package battlenet

// Client defines the Battle.net API client interface.
type Client interface {
	// UserAgent returns the client User-Agent header used in API requests.
	UserAgent() string
}

package client

import (
	"github.com/munsy/gobattlenet/pkg/quota"
)

type Response interface {
	// Body returns the response data as a byte slice.
	Body() []byte

	// Data returns the response as its appropriate type.
	Data() interface{}

	// Endpoint returns the endpoint that was used to obtain this response.
	Endpoint() string

	// Quota returns the current Battle.net API quota.
	Quota() quota.Quota
}

/*
// Response represents a Battle.net Account API response.
type Response struct {
	body []byte
	data interface{}
	endpoint string
	quota quota.Quota
}

// Body returns the response data as a byte slice.
func (r *Rseponse) Body() []byte {

}

	// Data returns the response as its appropriate type.
func (r *Rseponse) Data() interface{} {

}

	// Endpoint returns the endpoint that was used to obtain this response.
func (r *Rseponse) Endpoint() string {

}

	// Quota returns the current Battle.net API quota.
func (r *Rseponse) Quota() quota.Quota {

}

*/

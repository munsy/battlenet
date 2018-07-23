package wow

import (
	"github.com/munsy/gobattlenet/pkg/quota"
)

// Response represents a World of Warcraft API response.
type Response struct {
	Data     interface{}
	Endpoint string
	Quota    *quota.Quota
}

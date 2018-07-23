package sc2

import (
	"github.com/munsy/gobattlenet/pkg/quota"
)

// Response represents a Starcraft II API response.
type Response struct {
	Data     interface{}
	Endpoint string
	Quota    *quota.Quota
}

package d3

import (
	"github.com/munsy/gobattlenet/pkg/quota"
)

// Response represents a Diablo III API response.
type Response struct {
	Data     interface{}
	Endpoint string
	Quota    *quota.Quota
}

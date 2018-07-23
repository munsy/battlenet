package account

import (
	"github.com/munsy/gobattlenet/pkg/quota"
)

// Response represents a Battle.net Account API response.
type Response struct {
	Data     interface{}
	Endpoint string
	Quota    *quota.Quota
}

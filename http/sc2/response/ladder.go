package response

import (
	"github.com/munsy/battlenet/models/sc2"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// Ladder represents a Ladder response.
type Ladder struct {
	Data     *sc2.Ladder
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

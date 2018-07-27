package response

import (
	"github.com/munsy/battlenet/models/sc2"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// Rewards represents a Rewards response.
type Rewards struct {
	Data     *sc2.RewardsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

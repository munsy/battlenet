package response

import (
	"github.com/munsy/battlenet/models/sc2"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// Sc2Character represents a Sc2Character response.
type Sc2Character struct {
	Data     *sc2.Character
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

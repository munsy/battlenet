package response

import (
	"github.com/munsy/gobattlenet/models/sc2"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

// Sc2Character represents a Sc2Character response.
type Sc2Character struct {
	Data     *sc2.Character
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

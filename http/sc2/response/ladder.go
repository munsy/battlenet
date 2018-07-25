package response

import (
	"github.com/munsy/gobattlenet/models/sc2"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

type Ladder struct {
	Data     *sc2.Ladder
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

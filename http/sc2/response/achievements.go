package response

import (
	"github.com/munsy/gobattlenet/models/sc2"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

type Achievements struct {
	Data     *sc2.AchievementsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

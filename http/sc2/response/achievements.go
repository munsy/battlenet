package response

import (
	"github.com/munsy/gobattlenet/models/sc2"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

// Achievements represents an Achievements response.
type Achievements struct {
	Data     *sc2.AchievementsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

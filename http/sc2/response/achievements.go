package response

import (
	"github.com/munsy/battlenet/models/sc2"
	"github.com/munsy/battlenet/quota"
	"github.com/munsy/battlenet/regions"
)

// Achievements represents an Achievements response.
type Achievements struct {
	Data     *sc2.AchievementsData
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

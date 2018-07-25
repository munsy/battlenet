package response

import (
	"github.com/munsy/gobattlenet/models/account"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

type BattleID struct {
	Data     *account.BattleID
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

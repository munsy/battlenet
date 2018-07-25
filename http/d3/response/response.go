package response

import (
	"github.com/munsy/gobattlenet/models/d3"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

type ActIndex struct {
	Data     *d3.ActIndex
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type Act struct {
	Data     *d3.Act
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type Artisan struct {
	Data     *d3.Artisan
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type Recipe struct {
	Data     *d3.Recipe
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type Follower struct {
	Data     *d3.Follower
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type CharacterClass struct {
	Data     *d3.CharacterClass
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type CharacterSkill struct {
	Data     *d3.CharacterAPISkill
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type ItemTypeIndex struct {
	Data     *d3.ItemTypeIndex
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type ItemType struct {
	Data     *d3.ItemType
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type Item struct {
	Data     *d3.Item
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type Account struct {
	Data     *d3.Account
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type Hero struct {
	Data     *d3.Hero
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type HeroItems struct {
	Data     *d3.HeroItems
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

type HeroFollowers struct {
	Data     *d3.HeroFollowers
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

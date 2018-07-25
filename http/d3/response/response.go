package response

import (
	"github.com/munsy/gobattlenet/models/d3"
	"github.com/munsy/gobattlenet/quota"
	"github.com/munsy/gobattlenet/regions"
)

// ActIndex represents an ActIndex response.
type ActIndex struct {
	Data     *d3.ActIndex
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Act represents an Act response.
type Act struct {
	Data     *d3.Act
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Artisan represents an Artisan response.
type Artisan struct {
	Data     *d3.Artisan
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Recipe represents a Recipe response.
type Recipe struct {
	Data     *d3.Recipe
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Follower represents a Follower response.
type Follower struct {
	Data     *d3.Follower
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// CharacterClass represents a CharacterClass response.
type CharacterClass struct {
	Data     *d3.CharacterClass
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// CharacterSkill represents a CharacterSkill response.
type CharacterSkill struct {
	Data     *d3.CharacterAPISkill
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// ItemTypeIndex represents an ItemTypeIndex response.
type ItemTypeIndex struct {
	Data     *d3.ItemTypeIndex
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// ItemType represents an ItemType response.
type ItemType struct {
	Data     *d3.ItemType
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Item represents an Item response.
type Item struct {
	Data     *d3.Item
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Account represents an Account response.
type Account struct {
	Data     *d3.Account
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// Hero represents a Hero response.
type Hero struct {
	Data     *d3.Hero
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// HeroItems represents a HeroItems response.
type HeroItems struct {
	Data     *d3.HeroItems
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

// HeroFollowers represents a HeroFollowers response.
type HeroFollowers struct {
	Data     *d3.HeroFollowers
	Endpoint string
	Quota    *quota.Quota
	Region   regions.Region
}

package d3

/*
	Diablo III related API methods should go in here.
*/

// ActIndex gets an index of acts.
func (c *D3Client) ActIndex() (*ActIndex, error) {
	var index *ActIndex

	err := c.get(acts(c.region), index)

	if nil != err {
		return nil, err
	}

	return index, nil
}

// Act gets a single act by id.
func (c *D3Client) Act(actID int) (*Act, error) {
	var act *Act

	err := c.get(act(c.region, actID), act)

	if nil != err {
		return nil, err
	}

	return act, nil
}

// Artisan gets a single artisan by slug.
func (c *D3Client) Artisan(artisanSlug string) (*Artisan, error) {
	var artisan *Artisan

	err := c.get(artisan(c.region, artisanSlug), artisan)

	if nil != err {
		return nil, err
	}

	return artisan, nil
}

// Recipe gets a single recipe by slug for the specified artisan.
func (c *D3Client) Recipe(artisanSlug, recipeSlug string) (*Recipe, error) {
	var recipe *Recipe

	err := c.get(recipe(c.region, artisanSlug, recipeSlug), recipe)

	if nil != err {
		return nil, err
	}

	return recipe, nil
}

// Follower gets a single follower by slug.
func (c *D3Client) Follower(followerSlug string) (*Follower, error) {
	var follower *Follower

	err := c.get(follower(c.region, followerSlug), follower)

	if nil != err {
		return nil, err
	}

	return follower, nil
}

// CharacterClass gets a single character class by slug.
func (c *D3Client) CharacterClass(classSlug string) (*CharacterClass, error) {
	var cc *CharacterClass

	err := c.get(characterClass(c.region, classSlug), cc)

	if nil != err {
		return nil, err
	}

	return cc, nil
}

// CharacterSkill gets a single skill by slug, for a specific character class.
func (c *D3Client) CharacterSkill(classSlug, skillSlug string) (*CharacterAPISkill, error) {
	var cs *CharacterAPISkill

	err := c.get(skill(c.region, classSlug, skillSlug), cs)

	if nil != err {
		return nil, err
	}

	return cs, nil
}

// ItemTypeIndex gets an index of item types.
func (c *D3Client) ItemTypeIndex() (*ItemTypeIndex, error) {
	var index *ItemTypeIndex

	err := c.get(itemTypeIndex(c.region), index)

	if nil != err {
		return nil, err
	}

	return index, nil
}

// ItemType gets a single item type by slug.
func (c *D3Client) ItemType(itemTypeSlug string) (*ItemType, error) {
	var it *ItemType

	err := c.get(itemType(c.region, itemTypeSlug), it)

	if nil != err {
		return nil, err
	}

	return it, nil
}

// Item gets a single item by item slug and ID.
func (c *D3Client) Item(itemSlugAndID string) (*Item, error) {
	var item *Item

	err := c.get(item(c.region, itemSlugAndID), item)

	if nil != err {
		return nil, err
	}

	return item, nil
}

// Account gets a single profile.
func (c *D3Client) Account(account string) (*Account, error) {
	var acct *Account

	err := c.get(account(c.region, account), acct)

	if nil != err {
		return nil, err
	}

	return acct, nil
}

// Hero gets a single hero.
func (c *D3Client) Hero(account string, heroID int) (*Hero, error) {
	var hero *Hero

	err := c.get(hero(c.region, account, heroID), hero)

	if nil != err {
		return nil, err
	}

	return hero, nil
}

// HeroItems gets a list of items for the specified hero.
func (c *D3Client) HeroItems(account string, heroID int) (*HeroItems, error) {
	var items *HeroItems

	err := c.get(detailedHeroItems(c.region, account, heroID), items)

	if nil != err {
		return nil, err
	}

	return items, nil
}

// FollowerItems gets a list of items for the specified hero's followers.
func (c *D3Client) FollowerItems(account string, heroID int) (*HeroFollowers, error) {
	var followers *HeroFollowers

	err := c.get(detailedFollowerItems(c.region, account, heroID), followers)

	if nil != err {
		return nil, err
	}

	return followers, nil
}

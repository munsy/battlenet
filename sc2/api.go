package sc2

/*
	Starcraft related API methods should go in here.
*/

// Character returns the Sc2 character profile.
func (c *SC2Client) Character(name string, profileID int) (*Character, error) {
	var character *Character

	err := c.get(endpointProfile(c.region, profileID, c.region.Itoa(), name), &character)

	if nil != err {
		return nil, err
	}

	return character, nil
}

// LadderSeasons returns the Sc2 profile's ladder seasons.
func (c *SC2Client) LadderSeasons(name string, profileID int) (*LadderSeasons, error) {
	var ladderSeasons *LadderSeasons

	err := c.get(endpointLadderProfile(c.region, profileID, c.region.Itoa(), name), &ladderSeasons)

	if nil != err {
		return nil, err
	}

	return ladderSeasons, nil
}

// MatchHistory returns the Sc2 profile's match history.
func (c *SC2Client) MatchHistory(name string, profileID int) (*MatchHistory, error) {
	var matchHistory *MatchHistory

	err := c.get(endpointMatchHistory(c.region, profileID, c.region.Itoa(), name), &matchHistory)

	if nil != err {
		return nil, err
	}

	return matchHistory, nil
}

// Ladder returns Sc2 ladder data.
func (c *SC2Client) Ladder(id int) (*Ladder, error) {
	var ladder *Ladder

	err := c.get(endpointLadder(c.region, id), &ladder)

	if nil != err {
		return nil, err
	}

	return ladder, nil
}

// Achievements returns Sc2 achievement data.
func (c *SC2Client) Achievements(id int) (*AchievementsData, error) {
	var achievements *AchievementsData

	err := c.get(endpointAchievements(c.region), &achievements)

	if nil != err {
		return nil, err
	}

	return achievements, nil
}

// Rewards returns Sc2 reward data.
func (c *SC2Client) Rewards(id int) (*RewardsData, error) {
	var rewards *RewardsData

	err := c.get(endpointRewards(c.region), &rewards)

	if nil != err {
		return nil, err
	}

	return rewards, nil
}

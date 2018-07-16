package sc2

/*
	Starcraft related API methods should go in here.
*/

// Profile returns the Sc2 Profile.
func (c *SC2Client) Character(region, name string, profileID int) (*Character, error) {
	var character *Character

	err := c.get(profile(c.region, profileID, c.region.Int(), name), character)

	if nil != err {
		return nil, err
	}

	return character, nil
}

// LadderSeasons returns the Sc2 profile's ladder seasons.
func (c *SC2Client) LadderSeasons(region, name string, profileID int) (*LadderSeasons, error) {
	var ladderSeasons *LadderSeasons

	err := c.get(ladderProfile(c.region, profileID, c.region.Int(), name), ladderSeasons)

	if nil != err {
		return nil, err
	}

	return ladderSeasons, nil
}

// MatchHistory returns the Sc2 profile's match history.
func (c *SC2Client) MatchHistory(region, name string, profileID int) (*MatchHistory, error) {
	var matchHistory *MatchHistory

	err := c.get(matchHistory(c.region, profileID, c.region.Int(), name), matchHistory)

	if nil != err {
		return nil, err
	}

	return matchHistory, nil
}

// Ladder returns Sc2 ladder data.
func (c *SC2Client) Ladder(id int) (*Ladder, error) {
	var ladder *Ladder

	err := c.get(ladder(c.region, id), ladder)

	if nil != err {
		return nil, err
	}

	return ladder, nil
}

// Ladder returns Sc2 ladder data.
func (c *SC2Client) Achievements(id int) (*AchievementsData, error) {
	var achievements *AchievementsData

	err := c.get(achievements(c.region), achievements)

	if nil != err {
		return nil, err
	}

	return achievements, nil
}

// Ladder returns Sc2 ladder data.
func (c *SC2Client) Rewards(id int) (*RewardsData, error) {
	var rewards *RewardsData

	err := c.get(rewards(c.region), rewards)

	if nil != err {
		return nil, err
	}

	return rewards, nil
}

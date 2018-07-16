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

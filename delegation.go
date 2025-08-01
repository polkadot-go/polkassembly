package polkassembly

import "fmt"

func (c *Client) GetDelegationStats() (*DelegationStats, error) {
	var resp DelegationStats
	r, err := c.client.R().
		Get("/delegation/stats")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetDelegates(page, limit int) ([]Delegate, error) {
	var resp []Delegate
	queryParams := make(map[string]string)
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/delegation/delegates")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreatePADelegate(req CreatePADelegateRequest) (*Delegate, error) {
	var resp Delegate
	r, err := c.client.R().
		SetBody(req).
		Post("/delegation/pa-delegate")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdatePADelegate(delegateID int, req UpdatePADelegateRequest) (*Delegate, error) {
	var resp Delegate
	r, err := c.client.R().
		SetBody(req).
		Patch(fmt.Sprintf("/delegation/pa-delegate/%d", delegateID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPADelegate(delegateID int) (*Delegate, error) {
	var resp Delegate
	r, err := c.client.R().
		Get(fmt.Sprintf("/delegation/pa-delegate/%d", delegateID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeletePADelegate(delegateID int) error {
	r, err := c.client.R().
		Delete(fmt.Sprintf("/delegation/pa-delegate/%d", delegateID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) GetUserAllTracksStats(userID int) ([]TrackStats, error) {
	var resp []TrackStats
	r, err := c.client.R().
		Get(fmt.Sprintf("/delegation/users/%d/tracks/stats", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetUserTracksLevelData(userID int) ([]TrackLevelData, error) {
	var resp []TrackLevelData
	r, err := c.client.R().
		Get(fmt.Sprintf("/delegation/users/%d/tracks/levels", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

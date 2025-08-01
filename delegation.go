package polkassembly

import "fmt"

func (c *Client) GetDelegationStats() (*DelegationStats, error) {
	r, err := c.client.R().
		Get("/delegation/stats")
	if err != nil {
		return nil, err
	}
	var resp DelegationStats
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetDelegates(page, limit int) ([]Delegate, error) {
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
	var resp []Delegate
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreatePADelegate(req CreatePADelegateRequest) (*Delegate, error) {
	r, err := c.client.R().
		SetBody(req).
		Post("/delegation/delegates")
	if err != nil {
		return nil, err
	}
	var resp Delegate
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdatePADelegate(address string, manifesto string) (*Delegate, error) {
	r, err := c.client.R().
		SetBody(map[string]string{"manifesto": manifesto}).
		Patch(fmt.Sprintf("/delegation/delegates/%s", address))
	if err != nil {
		return nil, err
	}
	var resp Delegate
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPADelegate(address string) (*Delegate, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/delegation/delegates/%s", address))
	if err != nil {
		return nil, err
	}
	var resp Delegate
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeletePADelegate(address string) error {
	r, err := c.client.R().
		Delete(fmt.Sprintf("/delegation/delegates/%s", address))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) GetUserAllTracksStats(address string) ([]TrackStats, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/address/%s/delegation/tracks", address))
	if err != nil {
		return nil, err
	}
	var resp []TrackStats
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetUserTracksLevelData(address string, trackNum int) ([]TrackLevelData, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/address/%s/delegation/tracks/%d", address, trackNum))
	if err != nil {
		return nil, err
	}
	var resp []TrackLevelData
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

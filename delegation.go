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
		Post("/auth/actions/createPADelegate")
	if err != nil {
		return nil, err
	}
	var resp Delegate
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdatePADelegate(delegateID int, req UpdatePADelegateRequest) (*Delegate, error) {
	r, err := c.client.R().
		SetBody(req).
		Post(fmt.Sprintf("/auth/actions/updatePADelegate?delegateId=%d", delegateID))
	if err != nil {
		return nil, err
	}
	var resp Delegate
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPADelegate(delegateID int) (*Delegate, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/auth/query/paDelegate?delegateId=%d", delegateID))
	if err != nil {
		return nil, err
	}
	var resp Delegate
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeletePADelegate(delegateID int) error {
	r, err := c.client.R().
		Post(fmt.Sprintf("/auth/actions/deletePADelegate?delegateId=%d", delegateID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) GetUserAllTracksStats(userID int) ([]TrackStats, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/auth/query/userTracksStats?userId=%d", userID))
	if err != nil {
		return nil, err
	}
	var resp []TrackStats
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetUserTracksLevelData(userID int) ([]TrackLevelData, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/auth/query/userTracksLevel?userId=%d", userID))
	if err != nil {
		return nil, err
	}
	var resp []TrackLevelData
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

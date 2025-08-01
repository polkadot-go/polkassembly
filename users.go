package polkassembly

import "fmt"

func (c *Client) GetUserByID(userID int) (*User, error) {
	var resp User
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/%d", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUserFollowing(userID int, page, limit int) (*UserListingResponse, error) {
	var resp UserListingResponse
	queryParams := make(map[string]string)
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/users/%d/following", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUserFollowers(userID int, page, limit int) (*UserListingResponse, error) {
	var resp UserListingResponse
	queryParams := make(map[string]string)
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/users/%d/followers", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUserActivity(userID int, page, limit int) (*UserActivity, error) {
	var resp UserActivity
	queryParams := make(map[string]string)
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/users/%d/activity", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUserByUsername(username string) (*User, error) {
	var resp User
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/username/%s", username))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUserByAddress(address string) (*User, error) {
	var resp User
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/address/%s", address))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetUsers(params UserListingParams) (*UserListingResponse, error) {
	var resp UserListingResponse
	queryParams := make(map[string]string)
	if params.Page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.Limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", params.Limit)
	}
	if params.Sort != "" {
		queryParams["sort"] = params.Sort
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/users")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

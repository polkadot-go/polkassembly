package polkassembly

import (
	"fmt"
)

func (c *Client) GetUserByID(userID int) (*User, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/id/%d", userID))
	if err != nil {
		return nil, err
	}

	var resp User
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetUserFollowing(userID int, page, limit int) (*UserListingResponse, error) {
	queryParams := map[string]string{}
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/users/id/%d/following", userID))
	if err != nil {
		return nil, err
	}

	var resp UserListingResponse
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetUserFollowers(userID int, page, limit int) (*UserListingResponse, error) {
	queryParams := map[string]string{}
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/users/id/%d/followers", userID))
	if err != nil {
		return nil, err
	}

	var resp UserListingResponse
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetUserActivity(userID int, page, limit int) (*UserActivity, error) {
	queryParams := map[string]string{}
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/users/id/%d/activities", userID))
	if err != nil {
		return nil, err
	}

	var resp UserActivity
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetUserByUsername(username string) (*User, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/username/%s", username))
	if err != nil {
		return nil, err
	}

	var resp User
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetUserByAddress(address string) (*User, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/address/%s", address))
	if err != nil {
		return nil, err
	}

	var resp User
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetUsers(params UserListingParams) (*UserListingResponse, error) {
	queryParams := map[string]string{}
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

	var resp UserListingResponse
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

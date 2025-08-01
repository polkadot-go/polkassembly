package polkassembly

import "fmt"

func (c *Client) GetCartItems() ([]CartItem, error) {
	var resp []CartItem
	r, err := c.client.R().
		Get("/cart")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) AddCartItem(req AddCartItemRequest) (*CartItem, error) {
	var resp CartItem
	r, err := c.client.R().
		SetBody(req).
		Post("/cart")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateCartItem(itemID int, req UpdateCartItemRequest) (*CartItem, error) {
	var resp CartItem
	r, err := c.client.R().
		SetBody(req).
		Patch(fmt.Sprintf("/cart/%d", itemID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteCartItem(itemID int) error {
	r, err := c.client.R().
		Delete(fmt.Sprintf("/cart/%d", itemID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

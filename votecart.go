package polkassembly

import "fmt"

func (c *Client) GetCartItems() ([]CartItem, error) {
	var resp []CartItem

	// Auth endpoint, not a proposal type endpoint
	r, err := c.client.R().
		Get("/auth/query/cartItems")

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
		Post("/auth/actions/addCartItem")
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
		Post(fmt.Sprintf("/auth/actions/updateCartItem?itemId=%d", itemID))
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
		Post(fmt.Sprintf("/auth/actions/deleteCartItem?itemId=%d", itemID))
	if err != nil {
		return err
	}

	return c.parseResponse(r, nil)
}

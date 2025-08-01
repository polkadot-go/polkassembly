package polkassembly

import "fmt"

func (c *Client) GetCartItems(userID int) ([]CartItem, error) {
	var resp []CartItem
	r, err := c.client.R().
		Get(fmt.Sprintf("/users/id/%d/vote-cart", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) AddCartItem(userID int, req AddCartItemRequest) (*CartItem, error) {
	var resp CartItem
	r, err := c.client.R().
		SetBody(req).
		Post(fmt.Sprintf("/users/id/%d/vote-cart", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateCartItem(userID int, req UpdateCartItemRequest) (*CartItem, error) {
	var resp CartItem
	r, err := c.client.R().
		SetBody(req).
		Patch(fmt.Sprintf("/users/id/%d/vote-cart", userID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteCartItem(userID int, itemID string) error {
	r, err := c.client.R().
		SetBody(map[string]string{"id": itemID}).
		Delete(fmt.Sprintf("/users/id/%d/vote-cart", userID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

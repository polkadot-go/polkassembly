package polkassembly

import "fmt"

func (c *Client) GetPreimageForPost(postID int) (*Preimage, error) {
	var resp Preimage
	r, err := c.client.R().
		Get(fmt.Sprintf("/posts/%d/preimage", postID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPreimages(params PreimageListingParams) (*PreimageListingResponse, error) {
	var resp PreimageListingResponse
	queryParams := make(map[string]string)
	if params.Page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.Limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", params.Limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/preimages")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPreimageByHash(hash string) (*Preimage, error) {
	var resp Preimage
	r, err := c.client.R().
		Get(fmt.Sprintf("/preimages/%s", hash))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

package polkassembly

import "fmt"

func (c *Client) GetPreimageForPost(postID int) (*Preimage, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/posts/preimage?postId=%d", postID))

	if err != nil {
		return nil, err
	}

	var resp Preimage
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetPreimages(params PreimageListingParams) (*PreimageListingResponse, error) {
	queryParams := make(map[string]string)
	if params.Page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.Limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", params.Limit)
	}
	if params.Status != "" {
		queryParams["status"] = params.Status
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/preimages/list")

	if err != nil {
		return nil, err
	}

	var resp PreimageListingResponse
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetPreimageByHash(hash string) (*Preimage, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/preimages/hash?hash=%s", hash))

	if err != nil {
		return nil, err
	}

	var resp Preimage
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

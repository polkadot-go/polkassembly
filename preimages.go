package polkassembly

import "fmt"

func (c *Client) GetPreimages(params PreimageListingParams) (*PreimageListingResponse, error) {
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

	var resp PreimageListingResponse
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetPreimageByHash(hash string) (*Preimage, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/preimages/%s", hash))
	if err != nil {
		return nil, err
	}

	var resp Preimage
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

package polkassembly

import "fmt"

func (c *Client) GetVotes(params VoteListingParams) (*VoteListingResponse, error) {
	var resp VoteListingResponse
	queryParams := make(map[string]string)
	if params.PostID > 0 {
		queryParams["postId"] = fmt.Sprintf("%d", params.PostID)
	}
	if params.Page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.Limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", params.Limit)
	}
	if params.Decision != "" {
		queryParams["voteType"] = params.Decision
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/votes")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetVotesByAddress(address string, page, limit int) (*VoteListingResponse, error) {
	var resp VoteListingResponse
	queryParams := make(map[string]string{
		"voterAddress": address,
	})
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/votes/address-votes")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetVotesByUserID(userID int, page, limit int) (*VoteListingResponse, error) {
	var resp VoteListingResponse
	queryParams := make(map[string]string{
		"userId": fmt.Sprintf("%d", userID),
	})
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/votes/user-votes")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetVotingCurve(postID int) ([]VotingCurveData, error) {
	var resp struct {
		Curve []VotingCurveData `json:"curve"`
	}
	r, err := c.client.R().
		Get(fmt.Sprintf("/votes/curve?postId=%d", postID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp.Curve, nil
}

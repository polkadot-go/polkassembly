package polkassembly

import (
	"fmt"
)

// GetVotes retrieves votes for a specific proposal
// The API expects: /api/v2/{proposalType}/{postId}/votes
func (c *Client) GetVotes(params VoteListingParams) (*VoteListingResponse, error) {
	return c.GetVotesByType(params, "ReferendumV2")
}

// GetVotesByType retrieves votes for a specific proposal type
func (c *Client) GetVotesByType(params VoteListingParams, proposalType string) (*VoteListingResponse, error) {
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	queryParams := make(map[string]string)
	if params.Page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.Limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", params.Limit)
	}
	if params.VoteType != "" {
		queryParams["voteType"] = params.VoteType
	}

	// If postID is provided, get votes for specific post
	endpoint := fmt.Sprintf("/%s/votes", proposalType)
	if params.PostID > 0 {
		endpoint = fmt.Sprintf("/%s/%d/votes", proposalType, params.PostID)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	var resp VoteListingResponse
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetVotesByAddress retrieves votes by a specific address
func (c *Client) GetVotesByAddress(address string, page, limit int) (*VoteListingResponse, error) {
	queryParams := map[string]string{
		"voterAddress": address,
	}
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/votes/address")

	if err != nil {
		return nil, err
	}

	var resp VoteListingResponse
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetVotesByUserID retrieves votes by a specific user ID
func (c *Client) GetVotesByUserID(userID int, page, limit int) (*VoteListingResponse, error) {
	queryParams := map[string]string{
		"userId": fmt.Sprintf("%d", userID),
	}
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/votes/user")

	if err != nil {
		return nil, err
	}

	var resp VoteListingResponse
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetVotingCurve retrieves voting curve data for a proposal
func (c *Client) GetVotingCurve(postID int) ([]VotingCurveData, error) {
	return c.GetVotingCurveByType(postID, "ReferendumV2")
}

// GetVotingCurveByType retrieves voting curve data for a specific proposal type
func (c *Client) GetVotingCurveByType(postID int, proposalType string) ([]VotingCurveData, error) {
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	r, err := c.client.R().
		Get(fmt.Sprintf("/%s/%d/voting-curve", proposalType, postID))

	if err != nil {
		return nil, err
	}

	var resp struct {
		Curve []VotingCurveData `json:"curve"`
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return resp.Curve, nil
}

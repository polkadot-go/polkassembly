package polkassembly

import (
	"fmt"
)

// GetVotes retrieves votes for a specific proposal
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
	if params.Decision != "" {
		queryParams["decision"] = params.Decision
	}

	endpoint := fmt.Sprintf("/%s/%d/votes", proposalType, params.PostID)

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
func (c *Client) GetVotesByAddress(proposalType string, postID int, address string, page, limit int) (*VoteListingResponse, error) {
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	queryParams := map[string]string{}
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/%s/%d/votes/user/address/%s", proposalType, postID, address))
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
func (c *Client) GetVotesByUserID(proposalType string, postID int, userID int, page, limit int) (*VoteListingResponse, error) {
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	queryParams := map[string]string{}
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/%s/%d/votes/user/id/%d", proposalType, postID, userID))
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
		Get(fmt.Sprintf("/%s/%d/vote-curves", proposalType, postID))
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

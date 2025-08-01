package polkassembly

import "fmt"

func (c *Client) GetPosts(params PostListingParams) (*PostListingResponse, error) {
	var resp PostListingResponse
	queryParams := make(map[string]string)
	if params.Page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.ListingLimit > 0 {
		queryParams["listingLimit"] = fmt.Sprintf("%d", params.ListingLimit)
	}
	if params.TrackNo > 0 {
		queryParams["trackNo"] = fmt.Sprintf("%d", params.TrackNo)
	}
	if params.TrackStatus != "" {
		queryParams["trackStatus"] = params.TrackStatus
	}
	if params.ProposalType != "" {
		queryParams["proposalType"] = params.ProposalType
	}
	if params.SortBy != "" {
		queryParams["sortBy"] = params.SortBy
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/listing/on-chain-posts")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPost(postID int) (*Post, error) {
	var resp Post
	r, err := c.client.R().
		Get(fmt.Sprintf("/posts/on-chain-post?postId=%d", postID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPostOnchainData(postID int) (*PostOnchainData, error) {
	var resp PostOnchainData
	r, err := c.client.R().
		Get(fmt.Sprintf("/posts/on-chain-post-info?postId=%d", postID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetContentSummary(postID int) (*ContentSummary, error) {
	var resp ContentSummary
	r, err := c.client.R().
		Get(fmt.Sprintf("/posts/getContentSummary?postId=%d", postID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetPostComments(postID int) ([]Comment, error) {
	var resp struct {
		Comments []Comment `json:"comments"`
	}
	r, err := c.client.R().
		Get(fmt.Sprintf("/posts/comments?postId=%d&postType=on_chain", postID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp.Comments, nil
}

func (c *Client) GetActivityFeed(page, limit int) ([]ActivityFeedItem, error) {
	var resp struct {
		Feed []ActivityFeedItem `json:"feed"`
	}
	queryParams := make(map[string]string)
	if page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", page)
	}
	if limit > 0 {
		queryParams["limit"] = fmt.Sprintf("%d", limit)
	}

	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get("/posts/getAllActivity")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp.Feed, nil
}

func (c *Client) CreateOffchainPost(req CreateOffchainPostRequest) (*Post, error) {
	var resp Post
	r, err := c.client.R().
		SetBody(req).
		Post("/auth/actions/createPost")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdatePost(postID int, req UpdatePostRequest) (*Post, error) {
	var resp Post
	r, err := c.client.R().
		SetBody(req).
		Post(fmt.Sprintf("/auth/actions/editPost?postId=%d", postID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) IsSubscribed(postID int) (*SubscriptionStatus, error) {
	var resp SubscriptionStatus
	r, err := c.client.R().
		Get(fmt.Sprintf("/auth/query/isPostSubscribed?postId=%d", postID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetChildBounties(bountyID int) ([]Bounty, error) {
	var resp struct {
		ChildBounties []Bounty `json:"child_bounties"`
	}
	r, err := c.client.R().
		Get(fmt.Sprintf("/posts/child_bounties?parentBountyId=%d", bountyID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return resp.ChildBounties, nil
}

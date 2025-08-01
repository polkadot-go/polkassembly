package polkassembly

import (
	"encoding/json"
	"fmt"
)

// GetPosts retrieves posts based on proposal type
// The API expects: /api/v2/{proposalType}
func (c *Client) GetPosts(params PostListingParams) (*PostListingResponse, error) {
	// Default to ReferendumV2 if no type specified
	proposalType := params.ProposalType
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	queryParams := make(map[string]string)
	if params.Page > 0 {
		queryParams["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.ListingLimit > 0 {
		queryParams["listingLimit"] = fmt.Sprintf("%d", params.ListingLimit)
	}
	if params.SortBy != "" {
		queryParams["sortBy"] = params.SortBy
	}
	if params.TrackNo > 0 {
		queryParams["trackNo"] = fmt.Sprintf("%d", params.TrackNo)
	}
	if params.TrackStatus != "" {
		queryParams["trackStatus"] = params.TrackStatus
	}

	// The API expects proposalType as the main path
	r, err := c.client.R().
		SetQueryParams(queryParams).
		Get(fmt.Sprintf("/%s", proposalType))
	if err != nil {
		return nil, err
	}

	// Parse directly as PostListingResponse
	var resp PostListingResponse
	if err := json.Unmarshal(r.Body(), &resp); err != nil {
		return nil, fmt.Errorf("unmarshal posts: %w", err)
	}

	// Map items to Posts for backward compatibility and set PostID
	resp.Posts = resp.Items
	resp.Count = len(resp.Items)

	// Map Index to PostID for compatibility
	for i := range resp.Posts {
		if resp.Posts[i].PostID == 0 && resp.Posts[i].Index > 0 {
			resp.Posts[i].PostID = resp.Posts[i].Index
		}
		// Set username from publicUser if available
		if resp.Posts[i].PublicUser != nil && resp.Posts[i].Username == "" {
			resp.Posts[i].Username = resp.Posts[i].PublicUser.Username
		}
		// Set counts from metrics
		if resp.Posts[i].CommentsCount == 0 {
			resp.Posts[i].CommentsCount = resp.Posts[i].Metrics.Comments
		}
		if resp.Posts[i].ReactionsCount == 0 {
			resp.Posts[i].ReactionsCount = resp.Posts[i].Metrics.Reactions.Like + resp.Posts[i].Metrics.Reactions.Dislike
		}
		// Set status from onChainInfo
		if resp.Posts[i].Status == "" && resp.Posts[i].OnChainInfo != nil {
			resp.Posts[i].Status = resp.Posts[i].OnChainInfo.Status
		}
	}

	return &resp, nil
}

// GetPost retrieves a single post by ID
// The API expects: /api/v2/{proposalType}/{postId}
func (c *Client) GetPost(postID int) (*Post, error) {
	return c.GetPostByType(postID, "ReferendumV2")
}

// GetPostByType retrieves a single post by ID and type
func (c *Client) GetPostByType(postID int, proposalType string) (*Post, error) {
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	r, err := c.client.R().
		Get(fmt.Sprintf("/%s/%d", proposalType, postID))
	if err != nil {
		return nil, err
	}

	// Parse directly - single post responses may not be wrapped
	var resp Post
	if err := json.Unmarshal(r.Body(), &resp); err != nil {
		// Try wrapped response
		var wrapped struct {
			Data Post `json:"data"`
		}
		if err := json.Unmarshal(r.Body(), &wrapped); err != nil {
			return nil, fmt.Errorf("unmarshal post: %w", err)
		}
		resp = wrapped.Data
	}

	// Set backward compatible fields
	if resp.PostID == 0 && resp.Index > 0 {
		resp.PostID = resp.Index
	}
	if resp.PublicUser != nil && resp.Username == "" {
		resp.Username = resp.PublicUser.Username
	}

	return &resp, nil
}

// GetPostOnchainData retrieves onchain data for a post
func (c *Client) GetPostOnchainData(postID int) (*PostOnchainData, error) {
	return c.GetPostOnchainDataByType(postID, "ReferendumV2")
}

// GetPostOnchainDataByType retrieves onchain data for a post by type
func (c *Client) GetPostOnchainDataByType(postID int, proposalType string) (*PostOnchainData, error) {
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	// Try the onchain info endpoint
	r, err := c.client.R().
		Get(fmt.Sprintf("/%s/%d/onchain-info", proposalType, postID))
	if err != nil {
		return nil, err
	}

	// Check if response is HTML (404 page)
	if r.Header().Get("Content-Type") == "text/html" || len(r.Body()) > 0 && r.Body()[0] == '<' {
		return nil, fmt.Errorf("onchain data not available for post %d", postID)
	}

	var resp PostOnchainData
	if err := json.Unmarshal(r.Body(), &resp); err != nil {
		return nil, fmt.Errorf("unmarshal onchain data: %w", err)
	}

	return &resp, nil
}

// GetPostComments retrieves comments for a post
func (c *Client) GetPostComments(postID int) ([]Comment, error) {
	return c.GetPostCommentsByType(postID, "ReferendumV2")
}

// GetPostCommentsByType retrieves comments for a post by type
func (c *Client) GetPostCommentsByType(postID int, proposalType string) ([]Comment, error) {
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	r, err := c.client.R().
		Get(fmt.Sprintf("/%s/%d/comments", proposalType, postID))
	if err != nil {
		return nil, err
	}

	// Try parsing as array first
	var comments []Comment
	if err := json.Unmarshal(r.Body(), &comments); err == nil {
		return comments, nil
	}

	// Try parsing as object with comments field
	var resp struct {
		Comments []Comment `json:"comments"`
	}
	if err := json.Unmarshal(r.Body(), &resp); err != nil {
		return nil, fmt.Errorf("unmarshal comments: %w", err)
	}

	return resp.Comments, nil
}

// GetContentSummary retrieves AI-generated summary for a post
func (c *Client) GetContentSummary(postID int) (*ContentSummary, error) {
	return c.GetContentSummaryByType(postID, "ReferendumV2")
}

// GetContentSummaryByType retrieves AI-generated summary for a post by type
func (c *Client) GetContentSummaryByType(postID int, proposalType string) (*ContentSummary, error) {
	if proposalType == "" {
		proposalType = "ReferendumV2"
	}

	r, err := c.client.R().
		Get(fmt.Sprintf("/%s/%d/summary", proposalType, postID))
	if err != nil {
		return nil, err
	}

	var resp ContentSummary
	if err := json.Unmarshal(r.Body(), &resp); err != nil {
		return nil, fmt.Errorf("unmarshal summary: %w", err)
	}

	return &resp, nil
}

// GetActivityFeed retrieves the activity feed
func (c *Client) GetActivityFeed(page, limit int) ([]ActivityFeedItem, error) {
	// Activity feed is likely just posts endpoint
	posts, err := c.GetPosts(PostListingParams{
		Page:         page,
		ListingLimit: limit,
		ProposalType: "ReferendumV2",
	})
	if err != nil {
		return nil, err
	}

	// Convert posts to activity items
	var items []ActivityFeedItem
	for _, post := range posts.Posts {
		postID := post.PostID
		if postID == 0 {
			postID = post.Index
		}
		items = append(items, ActivityFeedItem{
			ID:       fmt.Sprintf("%d", postID),
			Type:     "post",
			PostID:   postID,
			PostType: post.ProposalType,
			Username: post.Username,
			Network:  post.Network,
			Content:  post.Title,
		})
	}

	return items, nil
}

// CreateOffchainPost creates an offchain discussion post
func (c *Client) CreateOffchainPost(req CreateOffchainPostRequest) (*Post, error) {
	// Set default post type
	if req.PostType == "" {
		req.PostType = "Discussion"
	}

	r, err := c.client.R().
		SetBody(req).
		Post("/api/v1/auth/actions/createPost")
	if err != nil {
		return nil, err
	}

	var resp Post
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdatePost updates an existing post
func (c *Client) UpdatePost(postID int, req UpdatePostRequest) (*Post, error) {
	r, err := c.client.R().
		SetBody(req).
		Post(fmt.Sprintf("/api/v1/auth/actions/editPost?postId=%d", postID))
	if err != nil {
		return nil, err
	}

	var resp Post
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// IsSubscribed checks if user is subscribed to a post
func (c *Client) IsSubscribed(postID int) (*SubscriptionStatus, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/ReferendumV2/%d/subscription", postID))
	if err != nil {
		return nil, err
	}

	var resp SubscriptionStatus
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetChildBounties retrieves child bounties for a parent bounty
func (c *Client) GetChildBounties(bountyID int) ([]Bounty, error) {
	r, err := c.client.R().
		Get(fmt.Sprintf("/Bounty/%d/child-bounties", bountyID))
	if err != nil {
		return nil, err
	}

	var resp struct {
		ChildBounties []Bounty `json:"child_bounties"`
	}
	if err := json.Unmarshal(r.Body(), &resp); err != nil {
		return nil, fmt.Errorf("unmarshal bounties: %w", err)
	}

	return resp.ChildBounties, nil
}

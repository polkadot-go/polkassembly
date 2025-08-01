package polkassembly

import "fmt"

func (c *Client) AddComment(req AddCommentRequest) (*Comment, error) {
	var resp Comment
	endpoint := fmt.Sprintf("/%s/%d/comments", req.PostType, req.PostID)
	r, err := c.client.R().
		SetBody(map[string]interface{}{
			"content":   req.Content,
			"sentiment": req.Sentiment,
			"parentId":  req.ParentID,
		}).
		Post(endpoint)
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) SubscribeProposal(postID int) error {
	r, err := c.client.R().
		Post(fmt.Sprintf("/ReferendumV2/%d/subscribe", postID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) UnsubscribeProposal(postID int) error {
	r, err := c.client.R().
		Post(fmt.Sprintf("/ReferendumV2/%d/unsubscribe", postID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) AddReaction(req AddReactionRequest) (*Reaction, error) {
	var resp Reaction
	r, err := c.client.R().
		SetBody(req).
		Post("/auth/actions/postReaction")
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateComment(commentID int, req UpdateCommentRequest) (*Comment, error) {
	var resp Comment
	r, err := c.client.R().
		SetBody(req).
		Post(fmt.Sprintf("/auth/actions/editComment?commentId=%d", commentID))
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteReaction(reactionID int) error {
	r, err := c.client.R().
		Post(fmt.Sprintf("/auth/actions/deleteReaction?reactionId=%d", reactionID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) FollowUser(userID int) error {
	r, err := c.client.R().
		Post(fmt.Sprintf("/auth/actions/followUser?userId=%d", userID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) UnfollowUser(userID int) error {
	r, err := c.client.R().
		Post(fmt.Sprintf("/auth/actions/unfollowUser?userId=%d", userID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

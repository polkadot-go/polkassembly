package polkassembly

import "fmt"

func (c *Client) AddComment(proposalType string, postID int, req AddCommentRequest) (*Comment, error) {
	var resp Comment
	endpoint := fmt.Sprintf("/%s/%d/comments", proposalType, postID)
	body := map[string]interface{}{
		"content": req.Content,
	}
	if req.ParentID != "" {
		body["parentCommentId"] = req.ParentID
	}
	if req.Address != "" {
		body["address"] = req.Address
	}
	r, err := c.client.R().
		SetBody(body).
		Post(endpoint)
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) AddReaction(proposalType string, postID int, reaction string) (*Reaction, error) {
	var resp Reaction
	endpoint := fmt.Sprintf("/%s/%d/reactions", proposalType, postID)
	r, err := c.client.R().
		SetBody(map[string]interface{}{
			"reaction": reaction,
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

func (c *Client) UpdateComment(proposalType string, postID int, commentID string, content interface{}) (*Comment, error) {
	var resp Comment
	endpoint := fmt.Sprintf("/%s/%d/comments/%s", proposalType, postID, commentID)
	r, err := c.client.R().
		SetBody(map[string]interface{}{
			"content": content,
		}).
		Patch(endpoint)
	if err != nil {
		return nil, err
	}
	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteComment(proposalType string, postID int, commentID string) error {
	endpoint := fmt.Sprintf("/%s/%d/comments/%s", proposalType, postID, commentID)
	r, err := c.client.R().
		Delete(endpoint)
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) DeleteReaction(proposalType string, postID int, reactionID string) error {
	endpoint := fmt.Sprintf("/%s/%d/reactions/%s", proposalType, postID, reactionID)
	r, err := c.client.R().
		Delete(endpoint)
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) FollowUser(userID int) error {
	r, err := c.client.R().
		Post(fmt.Sprintf("/users/id/%d/followers", userID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) UnfollowUser(userID int) error {
	r, err := c.client.R().
		Delete(fmt.Sprintf("/users/id/%d/followers", userID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) SubscribeProposal(proposalType string, postID int) error {
	r, err := c.client.R().
		Post(fmt.Sprintf("/%s/%d/subscription", proposalType, postID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

func (c *Client) UnsubscribeProposal(proposalType string, postID int) error {
	r, err := c.client.R().
		Delete(fmt.Sprintf("/%s/%d/subscription", proposalType, postID))
	if err != nil {
		return err
	}
	return c.parseResponse(r, nil)
}

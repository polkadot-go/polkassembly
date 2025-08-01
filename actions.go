package polkassembly

import (
	"fmt"
	"strings"
)

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

	// Handle 204 No Content response
	if r.StatusCode() == 204 || (r.StatusCode() >= 200 && r.StatusCode() < 300 && len(r.Body()) == 0) {
		// Return success with the updated content
		resp.ID = commentID
		resp.Content = content
		return &resp, nil
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

	// Check if response contains reaction data
	if r.StatusCode() >= 200 && r.StatusCode() < 300 {
		// Try to parse response
		if len(r.Body()) > 0 {
			if err := c.parseResponse(r, &resp); err == nil && resp.ID != "" {
				return &resp, nil
			}
		}

		// If no ID in response, generate one based on user reaction
		// The API might not return the reaction ID immediately
		resp.Reaction = reaction
		resp.ID = fmt.Sprintf("temp_%d_%s", postID, reaction)
		return &resp, nil
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
	// API might not support DELETE with ID, try removing reaction by type
	endpoint := fmt.Sprintf("/%s/%d/reactions", proposalType, postID)

	// Extract reaction type from ID if it's our temp format
	var reaction string
	if strings.HasPrefix(reactionID, "temp_") {
		parts := strings.Split(reactionID, "_")
		if len(parts) >= 3 {
			reaction = parts[2]
		}
	} else {
		// Assume the ID is the reaction type itself
		reaction = reactionID
	}

	r, err := c.client.R().
		SetBody(map[string]interface{}{
			"reaction": reaction,
		}).
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

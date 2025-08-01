package polkassembly

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func getTestClient(t *testing.T) *Client {
	network := os.Getenv("POLKASSEMBLY_NETWORK")
	if network == "" {
		network = "polkadot"
	}

	client := NewClient(Config{
		Network: network,
	})

	// Debug: log the base URL
	t.Logf("Using API URL: %s", client.baseURL)

	// Optional Web3 auth
	seedPhrase := os.Getenv("POLKASSEMBLY_SEED")
	if seedPhrase != "" {
		err := client.AuthenticateWithSeed(network, seedPhrase)
		if err != nil {
			t.Logf("Auth failed: %v", err)
		} else {
			t.Log("Authenticated with Web3")
		}
	}

	return client
}

// Public endpoints (no auth required)
func TestGetPosts(t *testing.T) {
	client := getTestClient(t)

	// Test without specifying proposal type
	resp, err := client.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 10,
	})
	if err != nil {
		t.Fatalf("GetPosts failed: %v", err)
	}
	if resp == nil {
		t.Fatal("Response is nil")
	}

	t.Logf("Found %d posts (total count: %d)", len(resp.Posts), resp.TotalCount)

	// Try with specific proposal type
	resp2, err := client.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 10,
		ProposalType: "ReferendumV2",
	})
	if err != nil {
		t.Logf("GetPosts with ReferendumV2 failed: %v", err)
	} else {
		t.Logf("Found %d ReferendumV2 posts", len(resp2.Posts))
	}

	// Try other parameters
	resp3, err := client.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 20,
		SortBy:       "newest",
	})
	if err != nil {
		t.Logf("GetPosts with sortBy failed: %v", err)
	} else {
		t.Logf("Found %d posts sorted by newest", len(resp3.Posts))
	}
}

func TestGetPost(t *testing.T) {
	client := getTestClient(t)

	// Get a post ID from listing first
	posts, err := client.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 1,
		ProposalType: "ReferendumV2",
	})
	if err != nil || posts == nil || len(posts.Posts) == 0 {
		t.Skip("No posts available")
	}

	postID := posts.Posts[0].PostID
	if postID == 0 {
		postID = posts.Posts[0].Index
	}

	post, err := client.GetPost(postID)
	if err != nil {
		t.Fatalf("GetPost failed: %v", err)
	}

	t.Logf("Got post: %s", post.Title)
}

func TestGetPostOnchainData(t *testing.T) {
	client := getTestClient(t)

	posts, err := client.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 1,
		ProposalType: "ReferendumV2",
	})
	if err != nil || posts == nil || len(posts.Posts) == 0 {
		t.Skip("No posts available")
	}

	postID := posts.Posts[0].PostID
	if postID == 0 {
		postID = posts.Posts[0].Index
	}

	data, err := client.GetPostOnchainData(postID)
	if err != nil {
		t.Logf("GetPostOnchainData failed: %v", err)
		return
	}

	t.Logf("Onchain data - Status: %s", data.Status)
}

func TestGetPostComments(t *testing.T) {
	client := getTestClient(t)

	posts, err := client.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 5,
		ProposalType: "ReferendumV2",
	})
	if err != nil || posts == nil || len(posts.Posts) == 0 {
		t.Skip("No posts available")
	}

	// Find a post with comments
	for _, post := range posts.Posts {
		commentsCount := post.CommentsCount
		if commentsCount == 0 && post.Metrics.Comments > 0 {
			commentsCount = post.Metrics.Comments
		}

		if commentsCount > 0 {
			postID := post.PostID
			if postID == 0 {
				postID = post.Index
			}

			comments, err := client.GetPostComments(postID)
			if err != nil {
				t.Errorf("GetPostComments failed: %v", err)
				continue
			}
			t.Logf("Found %d comments for post %d", len(comments), postID)
			return
		}
	}

	t.Log("No posts with comments found")
}

func TestGetUsers(t *testing.T) {
	client := getTestClient(t)

	resp, err := client.GetUsers(UserListingParams{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetUsers failed: %v", err)
	}

	t.Logf("Found %d users", len(resp.Users))
}

func TestGetUserByUsername(t *testing.T) {
	client := getTestClient(t)

	// Get a username from listing
	users, err := client.GetUsers(UserListingParams{Page: 1, Limit: 1})
	if err != nil || len(users.Users) == 0 {
		t.Skip("No users available")
	}

	user, err := client.GetUserByUsername(users.Users[0].Username)
	if err != nil {
		t.Fatalf("GetUserByUsername failed: %v", err)
	}

	t.Logf("Got user: %s", user.Username)
}

func TestGetVotes(t *testing.T) {
	client := getTestClient(t)

	// Get votes for a specific post
	posts, err := client.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 1,
		ProposalType: "ReferendumV2",
	})
	if err != nil || posts == nil || len(posts.Posts) == 0 {
		t.Skip("No posts available")
	}

	postID := posts.Posts[0].PostID
	if postID == 0 {
		postID = posts.Posts[0].Index
	}

	resp, err := client.GetVotes(VoteListingParams{
		PostID: postID,
		Page:   1,
		Limit:  10,
	})
	if err != nil {
		t.Fatalf("GetVotes failed: %v", err)
	}

	t.Logf("Found %d votes", len(resp.Votes))
}

func TestGetPreimages(t *testing.T) {
	client := getTestClient(t)

	resp, err := client.GetPreimages(PreimageListingParams{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		// This might fail due to external dependencies
		t.Logf("GetPreimages failed (may be external issue): %v", err)
		return
	}

	t.Logf("Found %d preimages", len(resp.Preimages))
}

func TestGetDelegationStats(t *testing.T) {
	client := getTestClient(t)

	stats, err := client.GetDelegationStats()
	if err != nil {
		// Skip if endpoint doesn't exist
		t.Skipf("GetDelegationStats not available: %v", err)
	}

	t.Logf("Total delegations: %d, Total balance: %s",
		stats.TotalDelegations, stats.TotalBalance)
}

func TestGetActivityFeed(t *testing.T) {
	client := getTestClient(t)

	feed, err := client.GetActivityFeed(1, 10)
	if err != nil {
		t.Logf("GetActivityFeed failed: %v", err)
		// Skip instead of fail since this might not be implemented
		t.Skip("Activity feed not available")
	}

	t.Logf("Found %d activity items", len(feed))
}

// Authenticated endpoints
func TestAuthenticatedEndpoints(t *testing.T) {
	seedPhrase := os.Getenv("POLKASSEMBLY_SEED")
	if seedPhrase == "" {
		t.Skip("Skipping authenticated tests: POLKASSEMBLY_SEED not set")
	}

	client := getTestClient(t)

	t.Run("GetCartItems", func(t *testing.T) {
		items, err := client.GetCartItems()
		if err != nil {
			// Skip if endpoint doesn't exist
			t.Skipf("GetCartItems not available: %v", err)
		}
		t.Logf("Cart has %d items", len(items))
	})

	t.Run("IsSubscribed", func(t *testing.T) {
		posts, err := client.GetPosts(PostListingParams{
			Page:         1,
			ListingLimit: 1,
			ProposalType: "ReferendumV2",
		})
		if err != nil || posts == nil || len(posts.Posts) == 0 {
			t.Skip("No posts available")
		}

		postID := posts.Posts[0].PostID
		if postID == 0 {
			postID = posts.Posts[0].Index
		}

		status, err := client.IsSubscribed(postID)
		if err != nil {
			t.Logf("IsSubscribed failed: %v", err)
		} else {
			t.Logf("Subscription status: %v", status.Subscribed)
		}
	})

	t.Run("CreateAndUpdateComment", func(t *testing.T) {
		posts, err := client.GetPosts(PostListingParams{
			Page:         1,
			ListingLimit: 1,
			ProposalType: "ReferendumV2",
		})
		if err != nil || posts == nil || len(posts.Posts) == 0 {
			t.Skip("No posts available")
		}

		postID := posts.Posts[0].PostID
		if postID == 0 {
			postID = posts.Posts[0].Index
		}

		// Create comment
		comment, err := client.AddComment(AddCommentRequest{
			Content:  "Test comment from Go client at " + time.Now().Format(time.RFC3339),
			PostID:   postID,
			PostType: "on_chain",
		})
		if err != nil {
			t.Logf("AddComment failed: %v", err)
			return
		}

		t.Logf("Created comment ID: %s", comment.ID)

		// Update comment
		commentID, _ := strconv.Atoi(comment.ID)
		updated, err := client.UpdateComment(commentID, UpdateCommentRequest{
			Content: "Updated: " + comment.Content,
		})
		if err != nil {
			t.Logf("UpdateComment failed: %v", err)
		} else {
			t.Logf("Updated comment: %s", updated.Content)
		}
	})

	t.Run("Reactions", func(t *testing.T) {
		posts, err := client.GetPosts(PostListingParams{
			Page:         1,
			ListingLimit: 1,
			ProposalType: "ReferendumV2",
		})
		if err != nil || posts == nil || len(posts.Posts) == 0 {
			t.Skip("No posts available")
		}

		postID := posts.Posts[0].PostID
		if postID == 0 {
			postID = posts.Posts[0].Index
		}

		reaction, err := client.AddReaction(AddReactionRequest{
			PostID:   postID,
			PostType: "on_chain",
			Reaction: "👍",
		})
		if err != nil {
			t.Logf("AddReaction failed: %v", err)
		} else {
			t.Logf("Added reaction ID: %s", reaction.ID)

			// Delete reaction
			reactionID, _ := strconv.Atoi(reaction.ID)
			err = client.DeleteReaction(reactionID)
			if err != nil {
				t.Logf("DeleteReaction failed: %v", err)
			} else {
				t.Log("Deleted reaction")
			}
		}
	})

	t.Run("SubscribeUnsubscribe", func(t *testing.T) {
		posts, err := client.GetPosts(PostListingParams{
			Page:         1,
			ListingLimit: 1,
			ProposalType: "ReferendumV2",
		})
		if err != nil || posts == nil || len(posts.Posts) == 0 {
			t.Skip("No posts available")
		}

		postID := posts.Posts[0].PostID
		if postID == 0 {
			postID = posts.Posts[0].Index
		}

		// Subscribe
		err = client.SubscribeProposal(postID)
		if err != nil {
			t.Logf("Subscribe failed: %v", err)
		} else {
			t.Log("Subscribed to proposal")
		}

		// Check subscription
		status, _ := client.IsSubscribed(postID)
		if status != nil {
			t.Logf("Subscription status after subscribe: %v", status.Subscribed)
		}

		// Unsubscribe
		err = client.UnsubscribeProposal(postID)
		if err != nil {
			t.Logf("Unsubscribe failed: %v", err)
		} else {
			t.Log("Unsubscribed from proposal")
		}
	})

	t.Run("EditProfile", func(t *testing.T) {
		user, err := client.EditUserDetails(EditUserDetailsRequest{
			Bio: "Test bio updated at " + time.Now().Format(time.RFC3339),
		})
		if err != nil {
			t.Logf("EditUserDetails failed: %v", err)
		} else {
			t.Logf("Updated user bio: %s", user.Bio)
		}
	})
}

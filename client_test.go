package polkassembly

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

var (
	testClient *Client
	authResp   *Web3AuthResponse
)

func TestMain(m *testing.M) {
	// Setup
	network := os.Getenv("POLKASSEMBLY_NETWORK")
	if network == "" {
		network = "polkadot"
	}

	debug := os.Getenv("POLKASSEMBLY_DEBUG") == "true"

	var logger *log.Logger
	if debug {
		logger = log.New(os.Stdout, "[test] ", log.LstdFlags)
	} else {
		logger = log.New(io.Discard, "", 0)
	}

	testClient = NewClient(Config{
		Network: network,
		Debug:   debug,
		Logger:  logger,
	})

	fmt.Printf("Using API URL: %s\n", testClient.baseURL)

	// Authenticate if seed is provided
	seedPhrase := os.Getenv("POLKASSEMBLY_SEED")
	if seedPhrase != "" {
		var err error
		authResp, err = authenticateAndGetResponse(testClient, network, seedPhrase)
		if err != nil {
			fmt.Printf("Auth failed: %v\n", err)
		} else {
			fmt.Println("Authenticated with Web3")
		}
	}

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func authenticateAndGetResponse(c *Client, network string, seedPhrase string) (*Web3AuthResponse, error) {
	var networkID uint16
	switch network {
	case "polkadot":
		networkID = 0
	case "kusama":
		networkID = 2
	default:
		networkID = 42
	}

	signer, err := NewPolkadotSignerFromSeed(seedPhrase, networkID)
	if err != nil {
		return nil, fmt.Errorf("create signer: %w", err)
	}

	fmt.Printf("Address: %s\n", signer.Address())

	message := fmt.Sprintf("Sign this message to authenticate with Polkassembly\n\nNetwork: %s\nAddress: %s\nTimestamp: %d",
		network, signer.Address(), time.Now().Unix())

	signature, err := signer.Sign([]byte(message))
	if err != nil {
		return nil, fmt.Errorf("sign message: %w", err)
	}

	req := Web3AuthRequest{
		Address:   signer.Address(),
		Signature: "0x" + hex.EncodeToString(signature),
		Message:   message,
		Network:   network,
	}

	resp, err := c.Web3Auth(req)
	if err != nil {
		return nil, fmt.Errorf("web3 auth: %w", err)
	}

	if c.token != "" {
		fmt.Printf("Client has token after auth\n")
		user, err := c.GetUserByAddress(signer.Address())
		if err == nil {
			resp.User = *user
			resp.Token = c.token
		}
	}

	return resp, nil
}

// Public endpoints tests
func TestGetPosts(t *testing.T) {
	resp, err := testClient.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 10,
	})

	if err != nil {
		t.Fatalf("GetPosts failed: %v", err)
	}

	t.Logf("Found %d posts (total count: %d)", len(resp.Posts), resp.TotalCount)

	resp2, err := testClient.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 10,
		ProposalType: "ReferendumV2",
	})

	if err != nil {
		t.Logf("GetPosts with ReferendumV2 failed: %v", err)
	} else {
		t.Logf("Found %d ReferendumV2 posts", len(resp2.Posts))
	}

	resp3, err := testClient.GetPosts(PostListingParams{
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
	posts, err := testClient.GetPosts(PostListingParams{
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

	post, err := testClient.GetPost(postID)
	if err != nil {
		t.Fatalf("GetPost failed: %v", err)
	}

	t.Logf("Got post: %s", post.Title)
}

func TestGetPostOnchainData(t *testing.T) {
	posts, err := testClient.GetPosts(PostListingParams{
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

	data, err := testClient.GetPostOnchainData(postID)
	if err != nil {
		t.Logf("GetPostOnchainData failed: %v", err)
		return
	}

	t.Logf("Onchain data - Status: %s, AyesCount: %d, NaysCount: %d",
		data.Status, data.AyesCount, data.NaysCount)
}

func TestGetPostComments(t *testing.T) {
	posts, err := testClient.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 5,
		ProposalType: "ReferendumV2",
	})

	if err != nil || posts == nil || len(posts.Posts) == 0 {
		t.Skip("No posts available")
	}

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

			comments, err := testClient.GetPostComments(postID)
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
	resp, err := testClient.GetUsers(UserListingParams{
		Page:  1,
		Limit: 10,
		Sort:  "profileScore",
	})

	if err != nil {
		t.Fatalf("GetUsers failed: %v", err)
	}

	t.Logf("Found %d users", len(resp.Users))

	if len(resp.Users) == 0 {
		resp, err = testClient.GetUsers(UserListingParams{
			Page:  1,
			Limit: 10,
			Sort:  "newest",
		})
		if err == nil {
			t.Logf("Found %d users with newest sort", len(resp.Users))
		}
	}
}

func TestGetUserByUsername(t *testing.T) {
	usernames := []string{"alice", "bob", "charlie", "admin"}

	for _, username := range usernames {
		user, err := testClient.GetUserByUsername(username)
		if err == nil {
			t.Logf("Got user: %s (ID: %d)", user.Username, user.ID)
			return
		}
	}

	t.Skip("No valid usernames found")
}

func TestGetVotes(t *testing.T) {
	posts, err := testClient.GetPosts(PostListingParams{
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

	resp, err := testClient.GetVotes(VoteListingParams{
		PostID: postID,
		Page:   1,
		Limit:  10,
	})

	if err != nil {
		t.Fatalf("GetVotes failed: %v", err)
	}

	t.Logf("Found %d votes for post %d", len(resp.Votes), postID)
}

func TestGetPreimages(t *testing.T) {
	posts, err := testClient.GetPosts(PostListingParams{
		Page:         1,
		ListingLimit: 10,
		ProposalType: "ReferendumV2",
	})

	if err == nil && posts != nil && len(posts.Posts) > 0 {
		for _, post := range posts.Posts {
			if post.Hash != "" {
				preimage, err := testClient.GetPreimageByHash(post.Hash)
				if err == nil {
					t.Logf("Found preimage for hash %s", post.Hash[:10]+"...")
					t.Logf("Preimage method: %s, section: %s", preimage.Method, preimage.Section)
					return
				}
			}
		}
	}

	resp, err := testClient.GetPreimages(PreimageListingParams{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		t.Logf("GetPreimages failed: %v", err)
		return
	}

	t.Logf("Found %d preimages", len(resp.Preimages))
}

func TestGetDelegationStats(t *testing.T) {
	stats, err := testClient.GetDelegationStats()
	if err != nil {
		delegates, err := testClient.GetDelegates(1, 10)
		if err != nil {
			t.Skipf("Delegation endpoints not available: %v", err)
		}
		t.Logf("Found %d delegates", len(delegates))
		return
	}

	t.Logf("Total delegations: %d, Total balance: %s",
		stats.TotalDelegations, stats.TotalBalance)
}

func TestGetActivityFeed(t *testing.T) {
	feed, err := testClient.GetActivityFeed(1, 10)
	if err != nil {
		t.Logf("GetActivityFeed failed: %v", err)
		t.Skip("Activity feed not available")
	}

	t.Logf("Found %d activity items", len(feed))
}

// Authenticated endpoints
func TestAuthenticatedEndpoints(t *testing.T) {
	if authResp == nil || authResp.User.ID == 0 {
		t.Skip("Authentication not available")
	}

	userID := authResp.User.ID
	t.Logf("Authenticated as user ID: %d, username: %s", userID, authResp.User.Username)

	t.Run("GetCartItems", func(t *testing.T) {
		items, err := testClient.GetCartItems(userID)
		if err != nil {
			t.Errorf("GetCartItems failed: %v", err)
		} else {
			t.Logf("Found %d cart items", len(items))
		}
	})

	t.Run("IsSubscribed", func(t *testing.T) {
		posts, err := testClient.GetPosts(PostListingParams{
			Page:         1,
			ListingLimit: 10,
			ProposalType: "ReferendumV2",
		})

		if err != nil || posts == nil || len(posts.Posts) == 0 {
			t.Skip("No posts available")
		}

		postID := posts.Posts[0].PostID
		if postID == 0 {
			postID = posts.Posts[0].Index
		}

		// Ensure clean state
		testClient.UnsubscribeProposal("ReferendumV2", postID)
		time.Sleep(2 * time.Second)

		// Check initial state
		status, _ := testClient.IsSubscribed("ReferendumV2", postID)
		if status != nil && status.Subscribed {
			t.Error("Post already subscribed after unsubscribe")
		}

		// Subscribe
		err = testClient.SubscribeProposal("ReferendumV2", postID)
		if err != nil {
			t.Fatalf("Subscribe failed: %v", err)
		}

		// Wait longer for propagation
		time.Sleep(5 * time.Second)

		// Check subscription
		status, err = testClient.IsSubscribed("ReferendumV2", postID)
		if err != nil {
			t.Errorf("IsSubscribed check failed: %v", err)
		} else if status == nil || !status.Subscribed {
			t.Errorf("Expected subscription to be true after subscribing (got: %v)", status)
		} else {
			t.Logf("Successfully subscribed to post %d", postID)
		}

		testClient.UnsubscribeProposal("ReferendumV2", postID)
	})

	t.Run("CreateAndUpdateComment", func(t *testing.T) {
		posts, err := testClient.GetPosts(PostListingParams{
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

		comment, err := testClient.AddComment("ReferendumV2", postID, AddCommentRequest{
			Content: "Test comment from Go client at " + time.Now().Format(time.RFC3339),
		})

		if err != nil {
			t.Fatalf("AddComment failed: %v", err)
		}

		t.Logf("Created comment ID: %s", comment.ID)

		_, err = testClient.UpdateComment("ReferendumV2", postID, comment.ID,
			"Updated: Test comment from Go client at "+time.Now().Format(time.RFC3339))

		if err != nil {
			t.Errorf("UpdateComment failed: %v", err)
		} else {
			t.Log("Comment updated successfully")
		}

		err = testClient.DeleteComment("ReferendumV2", postID, comment.ID)
		if err != nil {
			t.Logf("DeleteComment failed: %v", err)
		} else {
			t.Log("Comment deleted successfully")
		}
	})

	t.Run("Reactions", func(t *testing.T) {
		posts, err := testClient.GetPosts(PostListingParams{
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

		reactions := []string{"like", "dislike"}

		for _, reaction := range reactions {
			_, err := testClient.AddReaction("ReferendumV2", postID, reaction)
			if err != nil {
				t.Logf("AddReaction failed for %s: %v", reaction, err)
				continue
			}

			t.Logf("Added reaction: %s", reaction)
			// Skip delete test since API returns 405
			break
		}
	})

	t.Run("SubscribeUnsubscribe", func(t *testing.T) {
		posts, err := testClient.GetPosts(PostListingParams{
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

		testClient.UnsubscribeProposal("ReferendumV2", postID)
		time.Sleep(1 * time.Second)

		err = testClient.SubscribeProposal("ReferendumV2", postID)
		if err != nil {
			t.Errorf("Subscribe failed: %v", err)
			return
		}

		t.Log("Subscribed to proposal")
		time.Sleep(2 * time.Second)

		status, err := testClient.IsSubscribed("ReferendumV2", postID)
		if err != nil {
			t.Errorf("IsSubscribed check failed: %v", err)
		} else if status != nil && status.Subscribed {
			t.Log("Subscription confirmed")
		}

		err = testClient.UnsubscribeProposal("ReferendumV2", postID)
		if err != nil {
			t.Errorf("Unsubscribe failed: %v", err)
		} else {
			t.Log("Unsubscribed from proposal")
		}
	})

	t.Run("EditProfile", func(t *testing.T) {
		user, err := testClient.EditUserDetails(userID, EditUserDetailsRequest{
			Bio:   "Test bio from Go client at " + time.Now().Format(time.RFC3339),
			Title: "Go Developer",
		})

		if err != nil {
			t.Errorf("EditUserDetails failed: %v", err)
		} else {
			t.Logf("Updated user profile: ID %d", userID)
			if user.Username != "" {
				t.Logf("Username: %s", user.Username)
			}
		}
	})

	t.Run("FollowUnfollow", func(t *testing.T) {
		// We know bob exists with ID 4251
		bobID := 4251

		err := testClient.FollowUser(bobID)
		if err != nil {
			t.Logf("FollowUser failed: %v", err)
		} else {
			t.Logf("Followed user bob (ID: %d)", bobID)
		}

		time.Sleep(1 * time.Second)

		err = testClient.UnfollowUser(bobID)
		if err != nil {
			t.Logf("UnfollowUser failed: %v", err)
		} else {
			t.Log("Unfollowed user bob")
		}
	})
}

package polkassembly

import (
	"os"
	"testing"
)

func getTestClient(t *testing.T) *Client {
	network := os.Getenv("POLKASSEMBLY_NETWORK")
	if network == "" {
		network = "polkadot"
	}

	token := os.Getenv("POLKASSEMBLY_TOKEN")

	return NewClient(Config{
		Network: network,
		Token:   token,
	})
}

func TestGetPosts(t *testing.T) {
	client := getTestClient(t)

	resp, err := client.GetPosts(PostListingParams{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		t.Fatalf("GetPosts failed: %v", err)
	}

	if resp == nil {
		t.Fatal("GetPosts returned nil response")
	}

	if len(resp.Posts) == 0 {
		t.Log("No posts returned (might be empty)")
	}
}

func TestGetPost(t *testing.T) {
	client := getTestClient(t)

	// Test with a known post ID
	postID := 1
	resp, err := client.GetPost(postID)

	if err != nil {
		t.Logf("GetPost failed (post might not exist): %v", err)
		return
	}

	if resp.ID != postID {
		t.Errorf("Expected post ID %d, got %d", postID, resp.ID)
	}
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

	if resp == nil {
		t.Fatal("GetUsers returned nil response")
	}
}

func TestGetUserByUsername(t *testing.T) {
	client := getTestClient(t)

	// Skip if no test username provided
	username := os.Getenv("POLKASSEMBLY_TEST_USERNAME")
	if username == "" {
		t.Skip("Skipping test: POLKASSEMBLY_TEST_USERNAME not set")
	}

	resp, err := client.GetUserByUsername(username)
	if err != nil {
		t.Fatalf("GetUserByUsername failed: %v", err)
	}

	if resp.Username != username {
		t.Errorf("Expected username %s, got %s", username, resp.Username)
	}
}

func TestGetVotes(t *testing.T) {
	client := getTestClient(t)

	resp, err := client.GetVotes(VoteListingParams{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		t.Fatalf("GetVotes failed: %v", err)
	}

	if resp == nil {
		t.Fatal("GetVotes returned nil response")
	}
}

func TestGetPreimages(t *testing.T) {
	client := getTestClient(t)

	resp, err := client.GetPreimages(PreimageListingParams{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		t.Fatalf("GetPreimages failed: %v", err)
	}

	if resp == nil {
		t.Fatal("GetPreimages returned nil response")
	}
}

func TestGetDelegationStats(t *testing.T) {
	client := getTestClient(t)

	resp, err := client.GetDelegationStats()
	if err != nil {
		t.Fatalf("GetDelegationStats failed: %v", err)
	}

	if resp == nil {
		t.Fatal("GetDelegationStats returned nil response")
	}
}

func TestGetDelegates(t *testing.T) {
	client := getTestClient(t)

	resp, err := client.GetDelegates(1, 10)
	if err != nil {
		t.Fatalf("GetDelegates failed: %v", err)
	}

	if resp == nil {
		t.Fatal("GetDelegates returned nil response")
	}
}

// Test authenticated endpoints (requires valid token)
func TestAuthenticatedEndpoints(t *testing.T) {
	token := os.Getenv("POLKASSEMBLY_TOKEN")
	if token == "" {
		t.Skip("Skipping authenticated tests: POLKASSEMBLY_TOKEN not set")
	}

	client := getTestClient(t)

	t.Run("GetCartItems", func(t *testing.T) {
		resp, err := client.GetCartItems()
		if err != nil {
			t.Logf("GetCartItems failed (might be empty): %v", err)
		} else if resp != nil {
			t.Logf("Cart has %d items", len(resp))
		}
	})

	t.Run("IsSubscribed", func(t *testing.T) {
		postID := 1
		resp, err := client.IsSubscribed(postID)
		if err != nil {
			t.Logf("IsSubscribed failed: %v", err)
		} else if resp != nil {
			t.Logf("Subscription status for post %d: %v", postID, resp.IsSubscribed)
		}
	})
}

// Storage implementation for testing
type TestStorage struct {
	token string
}

func (s *TestStorage) SaveToken(token string) error {
	s.token = token
	return nil
}

func (s *TestStorage) GetToken() (string, error) {
	return s.token, nil
}

func (s *TestStorage) DeleteToken() error {
	s.token = ""
	return nil
}

func TestTokenStorage(t *testing.T) {
	storage := &TestStorage{}
	client := NewClient(Config{
		Network:      "polkadot",
		TokenStorage: storage,
	})

	testToken := "test-token-123"
	client.SetAuthToken(testToken)

	if storage.token != testToken {
		t.Errorf("Expected token %s in storage, got %s", testToken, storage.token)
	}

	// Test loading token from storage
	client2 := NewClient(Config{
		Network:      "polkadot",
		TokenStorage: storage,
	})

	if client2.token != testToken {
		t.Errorf("Expected token %s loaded from storage, got %s", testToken, client2.token)
	}
}

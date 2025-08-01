# Polkassembly Go Client

A complete Go client library for the Polkassembly v2 API, providing access to governance proposals, voting data, user management, and more.

## Installation

```
go get github.com/polkadot-go/polkassembly
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    polkassembly "github.com/polkadot-go/polkassembly"
)

func main() {
    // Create client
    client := polkassembly.NewClient(polkassembly.Config{
        Network: "polkadot", // or "kusama", "moonbeam", etc.
    })
    
    // Get posts
    posts, err := client.GetPosts(polkassembly.PostListingParams{
        Page:         1,
        ListingLimit: 10,
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d posts\n", len(posts.Posts))
}
```

## Authentication

### Web3 Authentication

```go
authResp, err := client.Web3Auth(polkassembly.Web3AuthRequest{
    Address:   "your-wallet-address",
    Signature: "signature",
    Message:   "message-to-sign",
})
```

### Web2 Authentication

```go
// Login
authResp, err := client.Web2Login(polkassembly.Web2LoginRequest{
    Username: "username",
    Password: "password",
})

// Signup
authResp, err := client.Web2Signup(polkassembly.Web2SignupRequest{
    Username: "username",
    Email:    "email@example.com",
    Password: "password",
})
```

### Token Storage

Implement the `TokenStorage` interface to persist authentication tokens:

```go
type MyStorage struct {
    // your storage implementation
}

func (s *MyStorage) SaveToken(token string) error {
    // Save token to database/file/etc
    return nil
}

func (s *MyStorage) GetToken() (string, error) {
    // Retrieve token
    return "", nil
}

func (s *MyStorage) DeleteToken() error {
    // Delete token
    return nil
}

// Use with client
client := polkassembly.NewClient(polkassembly.Config{
    Network:      "polkadot",
    TokenStorage: &MyStorage{},
})
```

## API Methods

### Posts

| Method | Description | Parameters | Returns |
|--------|-------------|------------|---------|
| `GetPosts` | List posts | `PostListingParams` | `*PostListingResponse` |
| `GetPost` | Get single post | `postID int` | `*Post` |
| `GetPostOnchainData` | Get onchain data | `postID int` | `*PostOnchainData` |
| `GetContentSummary` | Get content summary | `postID int` | `*ContentSummary` |
| `GetPostComments` | Get post comments | `postID int` | `[]Comment` |
| `GetActivityFeed` | Get activity feed | `page, limit int` | `[]ActivityFeedItem` |
| `CreateOffchainPost` | Create offchain post | `CreateOffchainPostRequest` | `*Post` |
| `UpdatePost` | Update post | `postID int, UpdatePostRequest` | `*Post` |
| `IsSubscribed` | Check subscription | `postID int` | `*SubscriptionStatus` |
| `GetChildBounties` | Get child bounties | `bountyID int` | `[]Bounty` |

### Votes

| Method | Description | Parameters | Returns |
|--------|-------------|------------|---------|
| `GetVotes` | List votes | `VoteListingParams` | `*VoteListingResponse` |
| `GetVotesByAddress` | Get votes by address | `address string, page, limit int` | `*VoteListingResponse` |
| `GetVotesByUserID` | Get votes by user | `userID, page, limit int` | `*VoteListingResponse` |
| `GetVotingCurve` | Get voting curve | `postID int` | `[]VotingCurveData` |

### Actions

| Method | Description | Parameters | Returns |
|--------|-------------|------------|---------|
| `AddComment` | Add comment | `AddCommentRequest` | `*Comment` |
| `AddReaction` | Add reaction | `AddReactionRequest` | `*Reaction` |
| `UpdateComment` | Update comment | `commentID int, UpdateCommentRequest` | `*Comment` |
| `DeleteReaction` | Delete reaction | `reactionID int` | `error` |
| `FollowUser` | Follow user | `userID int` | `error` |
| `UnfollowUser` | Unfollow user | `userID int` | `error` |
| `SubscribeProposal` | Subscribe to proposal | `postID int` | `error` |
| `UnsubscribeProposal` | Unsubscribe from proposal | `postID int` | `error` |

### Users

| Method | Description | Parameters | Returns |
|--------|-------------|------------|---------|
| `GetUserByID` | Get user by ID | `userID int` | `*User` |
| `GetUserByUsername` | Get user by username | `username string` | `*User` |
| `GetUserByAddress` | Get user by address | `address string` | `*User` |
| `GetUsers` | List users | `UserListingParams` | `*UserListingResponse` |
| `GetUserFollowing` | Get user following | `userID, page, limit int` | `*UserListingResponse` |
| `GetUserFollowers` | Get user followers | `userID, page, limit int` | `*UserListingResponse` |
| `GetUserActivity` | Get user activity | `userID, page, limit int` | `*UserActivity` |
| `EditUserDetails` | Edit user details | `EditUserDetailsRequest` | `*User` |

### Preimages

| Method | Description | Parameters | Returns |
|--------|-------------|------------|---------|
| `GetPreimageForPost` | Get preimage for post | `postID int` | `*Preimage` |
| `GetPreimages` | List preimages | `PreimageListingParams` | `*PreimageListingResponse` |
| `GetPreimageByHash` | Get preimage by hash | `hash string` | `*Preimage` |

### Vote Cart

| Method | Description | Parameters | Returns |
|--------|-------------|------------|---------|
| `GetCartItems` | Get cart items | - | `[]CartItem` |
| `AddCartItem` | Add cart item | `AddCartItemRequest` | `*CartItem` |
| `UpdateCartItem` | Update cart item | `itemID int, UpdateCartItemRequest` | `*CartItem` |
| `DeleteCartItem` | Delete cart item | `itemID int` | `error` |

### Delegation

| Method | Description | Parameters | Returns |
|--------|-------------|------------|---------|
| `GetDelegationStats` | Get delegation stats | - | `*DelegationStats` |
| `GetDelegates` | List delegates | `page, limit int` | `[]Delegate` |
| `CreatePADelegate` | Create PA delegate | `CreatePADelegateRequest` | `*Delegate` |
| `UpdatePADelegate` | Update PA delegate | `delegateID int, UpdatePADelegateRequest` | `*Delegate` |
| `GetPADelegate` | Get PA delegate | `delegateID int` | `*Delegate` |
| `DeletePADelegate` | Delete PA delegate | `delegateID int` | `error` |
| `GetUserAllTracksStats` | Get user track stats | `userID int` | `[]TrackStats` |
| `GetUserTracksLevelData` | Get user track levels | `userID int` | `[]TrackLevelData` |

## Testing

Run tests with environment variables:

```
# Basic tests (no authentication required)
POLKASSEMBLY_NETWORK=polkadot go test

# Full tests (requires authentication)
POLKASSEMBLY_NETWORK=polkadot \
POLKASSEMBLY_TOKEN=your-auth-token \
POLKASSEMBLY_TEST_USERNAME=test-username \
go test -v
```

## Error Handling

The library returns structured errors:

```go
posts, err := client.GetPosts(params)
if err != nil {
    if apiErr, ok := err.(*polkassembly.APIError); ok {
        fmt.Printf("API Error: %s\n", apiErr.Message)
    } else {
        fmt.Printf("Request Error: %v\n", err)
    }
}
```

## License

MIT License
# Polkassembly Go Client

A complete Go client library for the Polkassembly v2 API, providing access to governance proposals, voting data, user management, and more.

## Installation

```bash
go get github.com/polkadot-go/polkassembly
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "github.com/polkadot-go/polkassembly"
)

func main() {
    client := polkassembly.NewClient(polkassembly.Config{
        Network: "polkadot",
    })
    
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
err := client.AuthenticateWithSeed("polkadot", "your seed phrase here")
```

### Web2 Authentication
```go
authResp, err := client.Web2Login(polkassembly.Web2LoginRequest{
    Username: "username",
    Password: "password",
})
```

## Examples

See the `/examples` directory for complete examples:

- `load_all_referendums.go` - Load and filter all referendums
- `get_comments.go` - Get comments on a referendum  
- `filter_comments_sentiment.go` - Filter comments by sentiment
- `track_voting.go` - Track voting progress on proposals
- `authenticated_operations.go` - Comment, react, and subscribe
- `search_filter_proposals.go` - Search and filter proposals

## Configuration

### Debug Logging
```go
client := polkassembly.NewClient(polkassembly.Config{
    Network: "polkadot",
    Debug:   true,
})
```

### Token Storage
Implement the `TokenStorage` interface to persist authentication tokens.

## API Coverage

### Posts & Proposals
✅ List posts/proposals | Get single post | Get onchain data | Get comments | Create/update posts

### Voting  
✅ List votes | Get votes by address/user | Get voting curve data

### Users
✅ Get user info | List users | Follow/unfollow | Edit profile

### Actions (Authenticated)
✅ Add/update/delete comments | Add reactions | Subscribe/unsubscribe

### Delegation
✅ Get delegation stats | Manage delegates | Track stats

## Testing

```bash
# Basic tests
go test

# With authentication
export POLKASSEMBLY_SEED="your seed phrase"
export POLKASSEMBLY_NETWORK="polkadot"
go test -v
```


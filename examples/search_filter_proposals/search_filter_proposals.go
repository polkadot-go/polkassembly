package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/polkadot-go/polkassembly"
)

func main() {
	client := polkassembly.NewClient(polkassembly.Config{
		Network: "polkadot",
	})

	// Get treasury proposals
	treasuryProps, err := client.GetPosts(polkassembly.PostListingParams{
		ProposalType: "TreasuryProposal",
		ListingLimit: 50,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Filter by keyword
	keyword := "infrastructure"
	var filtered []polkassembly.Post

	for _, post := range treasuryProps.Posts {
		if strings.Contains(strings.ToLower(post.Title), keyword) ||
			strings.Contains(strings.ToLower(post.Content), keyword) {
			filtered = append(filtered, post)
		}
	}

	fmt.Printf("Found %d proposals mentioning '%s'\n", len(filtered), keyword)

	// Get proposals by specific track
	trackProps, err := client.GetPosts(polkassembly.PostListingParams{
		ProposalType: "ReferendumV2",
		TrackNo:      1, // Root track
		ListingLimit: 20,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Track 1 proposals: %d\n", len(trackProps.Posts))
}

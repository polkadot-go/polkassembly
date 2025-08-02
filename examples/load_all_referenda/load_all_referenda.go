package main

import (
	"fmt"
	"log"

	"github.com/polkadot-go/polkassembly-api"
)

func main() {
	client := polkassembly.NewClient(polkassembly.Config{
		Network: "polkadot",
	})

	var allPosts []polkassembly.Post
	page := 1

	for {
		resp, err := client.GetPosts(polkassembly.PostListingParams{
			Page:         page,
			ListingLimit: 100,
			ProposalType: "ReferendumV2",
		})
		if err != nil {
			log.Fatal(err)
		}

		allPosts = append(allPosts, resp.Posts...)

		if len(resp.Posts) < 100 {
			break
		}
		page++
	}

	fmt.Printf("Loaded %d referendums\n", len(allPosts))

	// Filter active referendums
	for _, post := range allPosts {
		if post.Status == "Active" || post.Status == "Deciding" {
			fmt.Printf("Active: #%d - %s\n", post.Index, post.Title)
		}
	}
}

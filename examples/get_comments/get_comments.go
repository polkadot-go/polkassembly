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

	referendumID := 1234 // Replace with actual referendum ID

	// Get referendum details
	post, err := client.GetPost(referendumID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Referendum #%d: %s\n", post.Index, post.Title)
	fmt.Printf("Comments: %d\n", post.Metrics.Comments)

	// Get comments
	comments, err := client.GetPostComments(referendumID)
	if err != nil {
		log.Fatal(err)
	}

	// Display comments
	for _, comment := range comments {
		content := fmt.Sprintf("%v", comment.Content)
		fmt.Printf("\n%s (%s):\n%s\n", comment.Username,
			comment.CreatedAt.Format("2006-01-02"), content)

		// Show replies
		for _, reply := range comment.Replies {
			replyContent := fmt.Sprintf("%v", reply.Content)
			fmt.Printf("  └─ %s: %s\n", reply.Username, replyContent)
		}
	}
}

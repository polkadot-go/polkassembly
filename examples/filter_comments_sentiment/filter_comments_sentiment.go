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

	referendumID := 1234
	comments, err := client.GetPostComments(referendumID)
	if err != nil {
		log.Fatal(err)
	}

	var positive, negative, neutral []polkassembly.Comment

	for _, comment := range comments {

		// Simple sentiment analysis
		if comment.Sentiment > 0 {
			positive = append(positive, comment)
		} else if comment.Sentiment < 0 {
			negative = append(negative, comment)
		} else {
			neutral = append(neutral, comment)
		}
	}

	fmt.Printf("Positive comments: %d\n", len(positive))
	fmt.Printf("Negative comments: %d\n", len(negative))
	fmt.Printf("Neutral comments: %d\n", len(neutral))
}

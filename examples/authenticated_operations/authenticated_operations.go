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

	// Authenticate
	err := client.AuthenticateWithSeed("polkadot", "your seed phrase")
	if err != nil {
		log.Fatal(err)
	}

	referendumID := 1234

	// Add comment
	comment, err := client.AddComment("ReferendumV2", referendumID,
		polkassembly.AddCommentRequest{
			Content: "This is my comment on the proposal",
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added comment: %s\n", comment.ID)

	// Add reaction
	_, err = client.AddReaction("ReferendumV2", referendumID, "like")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Added like reaction")

	// Subscribe to updates
	err = client.SubscribeProposal("ReferendumV2", referendumID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscribed to proposal updates")
}

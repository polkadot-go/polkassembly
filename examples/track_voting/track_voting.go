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

	referendumID := 1234

	// Get onchain data
	data, err := client.GetPostOnchainData(referendumID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %s\n", data.Status)
	fmt.Printf("Ayes: %d votes (%s DOT)\n", data.AyesCount, data.SupportAmount)
	fmt.Printf("Nays: %d votes (%s DOT)\n", data.NaysCount, data.AgainstAmount)

	// Get voting curve
	curve, err := client.GetVotingCurve(referendumID)
	if err != nil {
		log.Fatal(err)
	}

	for _, point := range curve {
		fmt.Printf("Block %d: Aye=%s, Nay=%s, Support=%s%%\n",
			point.BlockNumber, point.AyeAmount, point.NayAmount, point.Support)
	}
}

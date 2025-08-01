package polkassembly

import (
	"encoding/hex"
	"fmt"
	"time"
)

// Web3SignRequest represents a signing request
type Web3SignRequest struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

// Web3SignResponse represents the response from signing
type Web3SignResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
}

// AuthenticateWithSigner authenticates using a signer
func (c *Client) AuthenticateWithSigner(network string, signer Signer) error {
	// Generate a message to sign
	message := fmt.Sprintf("Sign this message to authenticate with Polkassembly\n\nNetwork: %s\nAddress: %s\nTimestamp: %d",
		network, signer.Address(), time.Now().Unix())

	// Sign the message
	signature, err := signer.Sign([]byte(message))
	if err != nil {
		return fmt.Errorf("sign message: %w", err)
	}

	// Create auth request
	req := Web3AuthRequest{
		Address:   signer.Address(),
		Signature: "0x" + hex.EncodeToString(signature),
		Message:   message,
		Network:   network,
	}

	// Authenticate
	resp, err := c.Web3Auth(req)
	if err != nil {
		return fmt.Errorf("web3 auth: %w", err)
	}

	// Store the token
	if resp.Token != "" {
		c.SetAuthToken(resp.Token)
	}

	return nil
}

// AuthenticateWithSeed authenticates using a seed phrase
func (c *Client) AuthenticateWithSeed(network string, seedPhrase string) error {
	// Determine network ID for SS58 encoding
	var networkID uint16
	switch network {
	case "polkadot":
		networkID = 0
	case "kusama":
		networkID = 2
	default:
		networkID = 42
	}

	// Create signer
	signer, err := NewPolkadotSignerFromSeed(seedPhrase, networkID)
	if err != nil {
		return fmt.Errorf("create signer: %w", err)
	}

	return c.AuthenticateWithSigner(network, signer)
}

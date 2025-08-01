package polkassembly

import (
	"fmt"
	"strings"

	"github.com/ChainSafe/go-schnorrkel"
	"github.com/vedhavyas/go-subkey/v2"
	"github.com/vedhavyas/go-subkey/v2/sr25519"
)

// Signer interface for message signing
type Signer interface {
	Sign(message []byte) ([]byte, error)
	Address() string
}

// PolkadotSigner implements the Signer interface for Polkadot accounts
type PolkadotSigner struct {
	privateKey *schnorrkel.SecretKey
	publicKey  *schnorrkel.PublicKey
	address    string
}

// NewPolkadotSignerFromSeed creates a new Polkadot signer from a seed phrase
func NewPolkadotSignerFromSeed(seedPhrase string, network uint16) (*PolkadotSigner, error) {
	seedPhrase = strings.TrimSpace(seedPhrase)

	scheme := sr25519.Scheme{}
	uri := seedPhrase

	kp, err := subkey.DeriveKeyPair(scheme, uri)
	if err != nil {
		return nil, fmt.Errorf("derive keypair: %w", err)
	}

	secretBytes := kp.Seed()
	if len(secretBytes) != 32 {
		return nil, fmt.Errorf("unexpected secret key length: %d", len(secretBytes))
	}

	var miniSecret [32]byte
	copy(miniSecret[:], secretBytes)

	miniSecretKey, err := schnorrkel.NewMiniSecretKeyFromRaw(miniSecret)
	if err != nil {
		return nil, fmt.Errorf("create mini secret key: %w", err)
	}

	secretKey := miniSecretKey.ExpandEd25519()
	publicKey, err := secretKey.Public()
	if err != nil {
		return nil, fmt.Errorf("get public key: %w", err)
	}

	var ss58Format uint16
	switch network {
	case 0: // Polkadot
		ss58Format = 0
	case 2: // Kusama
		ss58Format = 2
	default:
		ss58Format = 42 // Generic substrate
	}

	address := kp.SS58Address(ss58Format)

	return &PolkadotSigner{
		privateKey: secretKey,
		publicKey:  publicKey,
		address:    address,
	}, nil
}

// Sign signs a message using sr25519
func (s *PolkadotSigner) Sign(message []byte) ([]byte, error) {
	transcript := schnorrkel.NewSigningContext([]byte("substrate"), message)
	sig, err := s.privateKey.Sign(transcript)
	if err != nil {
		return nil, fmt.Errorf("sign message: %w", err)
	}
	sigBytes := sig.Encode()
	return sigBytes[:], nil
}

// Address returns the SS58 encoded address
func (s *PolkadotSigner) Address() string {
	return s.address
}

package polkassembly

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type TokenStorage interface {
	SaveToken(token string) error
	GetToken() (string, error)
	DeleteToken() error
}

type Client struct {
	client       *resty.Client
	baseURL      string
	token        string
	network      string
	tokenStorage TokenStorage
}

type Config struct {
	BaseURL      string
	Network      string
	Token        string
	Timeout      time.Duration
	TokenStorage TokenStorage
}

func NewClient(cfg Config) *Client {
	if cfg.BaseURL == "" {
		// Use the correct API v2 base URL
		cfg.BaseURL = fmt.Sprintf("https://%s.polkassembly.io/api/v2", cfg.Network)
	}

	if cfg.Timeout == 0 {
		cfg.Timeout = 90 * time.Second
	}

	client := resty.New().
		SetBaseURL(cfg.BaseURL).
		SetTimeout(cfg.Timeout).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("x-network", cfg.Network)

	c := &Client{
		client:       client,
		baseURL:      cfg.BaseURL,
		network:      cfg.Network,
		token:        cfg.Token,
		tokenStorage: cfg.TokenStorage,
	}

	if cfg.Token != "" {
		c.SetAuthToken(cfg.Token)
	} else if cfg.TokenStorage != nil {
		if token, err := cfg.TokenStorage.GetToken(); err == nil && token != "" {
			c.SetAuthToken(token)
		}
	}

	return c
}

func (c *Client) SetAuthToken(token string) {
	c.token = token
	c.client.SetHeader("Authorization", "Bearer "+token)

	if c.tokenStorage != nil {
		c.tokenStorage.SaveToken(token)
	}
}

func (c *Client) SetNetwork(network string) {
	c.network = network
	c.client.SetHeader("x-network", network)
}

func (c *Client) parseResponse(resp *resty.Response, v interface{}) error {
	// Debug logging
	if resp.StatusCode() >= 400 {
		fmt.Printf("Error response: %d - %s\n", resp.StatusCode(), string(resp.Body()))
	}

	if resp.IsError() {
		var apiErr APIError
		if err := json.Unmarshal(resp.Body(), &apiErr); err != nil {
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		}
		return &apiErr
	}

	if v != nil && len(resp.Body()) > 0 {
		// Debug: log response for empty results
		var temp map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &temp); err == nil {
			if posts, ok := temp["posts"]; ok {
				if postsArr, ok := posts.([]interface{}); ok && len(postsArr) == 0 {
					fmt.Printf("Empty posts response from %s\n", resp.Request.URL)
				}
			}
		}

		if err := json.Unmarshal(resp.Body(), v); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
	}

	return nil
}

func (c *Client) handleAuthResponse(token string) {
	if token != "" {
		c.SetAuthToken(token)
	}
}

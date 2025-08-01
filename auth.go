package polkassembly

import "fmt"

func (c *Client) Web3Auth(req Web3AuthRequest) (*Web3AuthResponse, error) {
	var resp Web3AuthResponse

	if req.Network == "" {
		req.Network = c.network
	}

	r, err := c.client.R().
		SetBody(req).
		Post("/auth/web3-auth")

	if err != nil {
		return nil, err
	}

	c.logDebug("Auth response status: %d", r.StatusCode())

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	// Store all cookies from auth response
	for _, cookie := range r.Cookies() {
		c.client.SetCookie(cookie)
		c.logDebug("Storing cookie: %s", cookie.Name)

		if cookie.Name == "access_token" {
			resp.Token = cookie.Value
			c.SetAuthToken(cookie.Value)
		}
	}

	return &resp, nil
}

func (c *Client) Web2Login(req Web2LoginRequest) (*Web2LoginResponse, error) {
	var resp Web2LoginResponse

	r, err := c.client.R().
		SetBody(req).
		Post("/auth/web2-auth/login")

	if err != nil {
		return nil, err
	}

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	c.handleAuthResponse(resp.Token)
	return &resp, nil
}

func (c *Client) Web2Signup(req Web2SignupRequest) (*Web2SignupResponse, error) {
	var resp Web2SignupResponse

	r, err := c.client.R().
		SetBody(req).
		Post("/auth/web2-auth/signup")

	if err != nil {
		return nil, err
	}

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	c.handleAuthResponse(resp.Token)
	return &resp, nil
}

func (c *Client) SendResetPasswordEmail(req ResetPasswordRequest) error {
	r, err := c.client.R().
		SetBody(req).
		Post("/auth/send-reset-password-email")

	if err != nil {
		return err
	}

	return c.parseResponse(r, nil)
}

func (c *Client) ResetPasswordWithToken(token, newPassword string) error {
	r, err := c.client.R().
		SetBody(map[string]string{
			"token":       token,
			"newPassword": newPassword,
		}).
		Post("/auth/reset-password-with-token")

	if err != nil {
		return err
	}

	return c.parseResponse(r, nil)
}

func (c *Client) GenerateQRSession() (*QRSessionResponse, error) {
	var resp QRSessionResponse

	r, err := c.client.R().
		Get("/auth/qr-session")

	if err != nil {
		return nil, err
	}

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) ClaimQRSession(req ClaimQRSessionRequest) (*Web3AuthResponse, error) {
	var resp Web3AuthResponse

	r, err := c.client.R().
		SetBody(req).
		Post("/auth/qr-session")

	if err != nil {
		return nil, err
	}

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	c.handleAuthResponse(resp.Token)
	return &resp, nil
}

func (c *Client) EditUserDetails(userID int, req EditUserDetailsRequest) (*User, error) {
	var resp User

	r, err := c.client.R().
		SetBody(req).
		Patch(fmt.Sprintf("/users/id/%d", userID))

	if err != nil {
		return nil, err
	}

	if r.StatusCode() == 204 {
		return c.GetUserByID(userID)
	}

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

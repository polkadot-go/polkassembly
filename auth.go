package polkassembly

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

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	c.handleAuthResponse(resp.Token)
	return &resp, nil
}

func (c *Client) Web2Login(req Web2LoginRequest) (*Web2LoginResponse, error) {
	var resp Web2LoginResponse
	r, err := c.client.R().
		SetBody(req).
		Post("/auth/actions/login")
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
		Post("/auth/actions/signup")
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
		Post("/auth/actions/sendResetPasswordEmail")
	if err != nil {
		return err
	}

	return c.parseResponse(r, nil)
}

func (c *Client) GenerateQRSession() (*QRSessionResponse, error) {
	var resp QRSessionResponse
	r, err := c.client.R().
		Post("/auth/actions/qrGenerateSession")
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
		Post("/auth/actions/qrClaimSession")
	if err != nil {
		return nil, err
	}

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	c.handleAuthResponse(resp.Token)
	return &resp, nil
}

func (c *Client) EditUserDetails(req EditUserDetailsRequest) (*User, error) {
	var resp User
	r, err := c.client.R().
		SetBody(req).
		Post("/auth/actions/editProfile")
	if err != nil {
		return nil, err
	}

	if err := c.parseResponse(r, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

package pocket

type Client struct {
	ConsumerKey string
	Token       string
}

type RetrieveResult struct {
	List map[string]interface{}
}

type RetrieveOption struct {
	Search string `json:"search,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

type authInfo struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

func (c *Client) Get() (int, error) {
	consumerKey, err := getConsumerKey(c.ConsumerKey)
	if err != nil {
		return 0, err
	}
	token, err := getAccessToken(c.Token)
	if err != nil {
		return 0, err
	}

	data := authInfo{
		ConsumerKey: consumerKey,
		AccessToken: token,
	}

	res := &RetrieveResult{}
	err = PostJSON("/v3/get", data, res)
	if err != nil {
		return 0, err
	}

	return len(res.List), nil
}

package pocket

type Client struct {
	authInfo
}

type authInfo struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

type RetrieveOption struct {
	Search string `json:"search,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

type RetrieveResult struct {
	List     map[string]interface{}
	Status   int
	Complete int
	Since    int
}

type retrieveAPIOptionWithAuth struct {
	*RetrieveOption
	authInfo
}

func NewClient(consumerKey, accessToken string) *Client {
	return &Client{
		authInfo: authInfo{
			ConsumerKey: consumerKey,
			AccessToken: accessToken,
		},
	}
}

func (c *Client) Retrieve(options *RetrieveOption) (*RetrieveResult, error) {
	data := retrieveAPIOptionWithAuth{
		authInfo:       c.authInfo,
		RetrieveOption: options,
	}

	res := &RetrieveResult{}
	err := PostJSON("/v3/get", data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

package pocket

import (
	"context"
)

type Pocket struct {
	ConsumerKeyFile string
	TokenFile       string
}

func (p *Pocket) Get(ctx context.Context) (int, error) {
	consumerKey, err := getConsumerKey(p.ConsumerKeyFile)
	if err != nil {
		return 0, err
	}
	token, err := getAccessToken(p.TokenFile)
	if err != nil {
		return 0, err
	}
	client := NewClient(consumerKey, token)
	res, err := client.Retrieve(nil)
	if err != nil {
		return 0, err
	}
	return len(res.List), nil
}

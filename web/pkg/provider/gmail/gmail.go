package gmail

import (
	"context"

	"github.com/yuseinishiyama/stats/pkg/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Client struct {
	Credential string
	Token      string
}

func (g *Client) Get() (int, error) {
	ctx := context.Background()
	config, err := google.GetConfig(g.Credential)
	if err != nil {
		return 0, err
	}
	client, err := google.GetClient(ctx, config, g.Token)
	if err != nil {
		return 0, err
	}
	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return 0, err
	}
	r, err := srv.Users.Labels.Get("me", "INBOX").Do()
	if err != nil {
		return 0, err
	}
	return int(r.MessagesTotal), nil
}

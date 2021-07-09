package provider

import (
	"context"

	"github.com/yuseinishiyama/stats/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Gmail struct {
	Credential string
	Token      string
}

func (g *Gmail) Get(ctx context.Context) (int, error) {
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

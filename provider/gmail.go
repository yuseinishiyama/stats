package provider

import (
	"context"
	"log"

	"github.com/yuseinishiyama/stats/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Gmail struct {
	Credential string
	Token      string
}

func (g *Gmail) Get() int64 {
	ctx := context.Background()
	client := google.GetClient(ctx, g.Credential, g.Token)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Labels.Get(user, "INBOX").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}

	return r.MessagesTotal
}

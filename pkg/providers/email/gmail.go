package providers

import (
	"context"
	"log"

	"github.com/yuseinishiyama/stats/pkg/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func Run() int64 {
	ctx := context.Background()
	client := google.GetClient(ctx, "config/google-work-credentials.json", "config/google-work-token.json")

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

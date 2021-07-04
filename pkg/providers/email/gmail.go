package providers

import (
	"context"
	"fmt"
	"log"

	"github.com/yuseinishiyama/stats/pkg/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func Run() {
	ctx := context.Background()
	client := google.GetClient(ctx, "config/credentials.json", "config/token.json")

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Labels.Get(user, "INBOX").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	fmt.Println(r.MessagesTotal)
}

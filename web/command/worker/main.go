package worker

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/yuseinishiyama/stats/pkg/provider/gmail"
	"github.com/yuseinishiyama/stats/pkg/provider/pocket"
	"github.com/yuseinishiyama/stats/pkg/provider/slack"
	"github.com/yuseinishiyama/stats/pkg/storage"
)

type worker struct {
	spreadsheet *storage.Spreadsheet
	context     context.Context
}

func Command() *cobra.Command {
	worker := worker{}

	cmd := &cobra.Command{
		Use:   "worker",
		Short: "collects personal statistics",
		Run: func(cmd *cobra.Command, args []string) {
			worker.Execute()
		},
	}

	return cmd
}

func (r *worker) Execute() {
	r.context = context.Background()
	r.spreadsheet = &storage.Spreadsheet{
		ID:         "1yG-Hzw4_U4wnEZMUToNGxb7v8_-Ab60BJrgTk6T4798",
		Credential: "config/google-private-credential.json",
		Token:      "config/google-private-token.json",
	}

	if err := r.updateWorkGmail(); err != nil {
		log.Printf("Failed to update work gmail inbox count: %v", err)
	}
	if err := r.updatePrivateGmail(); err != nil {
		log.Printf("Failed to update private gmail inbox count: %v", err)
	}
	if err := r.updatePocket(); err != nil {
		log.Printf("Failed to update pocket unread count: %v", err)
	}
	if err := r.updateSlack(); err != nil {
		log.Printf("Failed to update slack saved item count: %v", err)
	}
}

func (r *worker) updateWorkGmail() error {
	client := &gmail.Client{
		Credential: "config/google-work-credential.json",
		Token:      "config/google-work-token.json",
	}
	val, err := client.Get()
	if err != nil {
		return err
	}
	mailInboxWork := storage.NewMailInboxWorkEntry(val)
	return r.spreadsheet.Write(r.context, mailInboxWork)
}

func (r *worker) updatePrivateGmail() error {
	client := &gmail.Client{
		Credential: "config/google-private-credential.json",
		Token:      "config/google-private-token.json",
	}
	val, err := client.Get()
	if err != nil {
		return err
	}
	mailInboxPrivate := storage.NewMailInboxPrivateEntry(val)
	return r.spreadsheet.Write(r.context, mailInboxPrivate)
}

func (r *worker) updatePocket() error {
	client := &pocket.Client{
		ConsumerKey: "config/pocket-consumer-key",
		Token:       "config/pocket-token.json",
	}
	val, err := client.Get()
	if err != nil {
		return err
	}
	mailInboxPrivate := storage.NewReadItLaterEntry(val)
	return r.spreadsheet.Write(r.context, mailInboxPrivate)
}

func (r *worker) updateSlack() error {
	client := &slack.Client{TokenFile: "config/slack-token"}
	val, err := client.Get()
	if err != nil {
		return err
	}
	mailInboxPrivate := storage.NewChatSavedEntry(val)
	return r.spreadsheet.Write(r.context, mailInboxPrivate)
}

package command

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/yuseinishiyama/stats/command/gentoken"
	"github.com/yuseinishiyama/stats/provider"
	"github.com/yuseinishiyama/stats/storage"
)

type rootCommand struct {
	spreadsheet *storage.Spreadsheet
	context     context.Context
}

func Command() *cobra.Command {
	rootCmd := rootCommand{}

	cmd := &cobra.Command{
		Use:   "stats",
		Short: "reports personal statistics",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.Execute()
		},
	}

	cmd.AddCommand(gentoken.Command())
	return cmd
}

func (r *rootCommand) Execute() {
	r.context = context.Background()
	r.spreadsheet = &storage.Spreadsheet{
		ID:         "1yG-Hzw4_U4wnEZMUToNGxb7v8_-Ab60BJrgTk6T4798",
		Credential: "config/google-private-credential.json",
		Token:      "config/google-private-token.json",
	}

	if err := r.updateWorkGmail(); err != nil {
		log.Fatalf("Failed to update work gmail inbox count: %v", err)
	}
	if err := r.updatePrivateGmail(); err != nil {
		log.Fatalf("Failed to update private gmail inbox count: %v", err)
	}
}

func (r *rootCommand) updateWorkGmail() error {
	workGmail := &provider.Gmail{
		Credential: "config/google-work-credential.json",
		Token:      "config/google-work-token.json",
	}
	val, err := workGmail.Get(r.context)
	if err != nil {
		return err
	}
	mailInboxWork := storage.NewMailInboxWorkEntry(*val)
	return r.spreadsheet.Write(r.context, mailInboxWork)
}

func (r *rootCommand) updatePrivateGmail() error {
	privateGmail := &provider.Gmail{
		Credential: "config/google-private-credential.json",
		Token:      "config/google-private-token.json",
	}
	val, err := privateGmail.Get(r.context)
	if err != nil {
		return err
	}
	mailInboxPrivate := storage.NewMailInboxPrivateEntry(*val)
	return r.spreadsheet.Write(r.context, mailInboxPrivate)
}

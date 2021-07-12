package bot

import (
	"log"

	"github.com/spf13/cobra"
)

type bot struct{}

func Command() *cobra.Command {
	bot := bot{}

	cmd := &cobra.Command{
		Use:   "bot",
		Short: "runs input bot",
		Run: func(cmd *cobra.Command, args []string) {
			bot.Execute()
		},
	}

	return cmd
}

func (i *bot) Execute() {
	log.Println("TODO")
}

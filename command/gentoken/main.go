package gentoken

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yuseinishiyama/stats/google"
)

type genToken struct{}

func Command() *cobra.Command {
	genToken := genToken{}

	cmd := &cobra.Command{
		Use:   "gen-token",
		Short: "generates tokens",
		Run: func(cmd *cobra.Command, args []string) {
			genToken.Execute()
		},
	}

	return cmd
}

func (g *genToken) Execute() {
	if err := google.GenerateToken("config/google-work-credential.json", "config/google-work-token.json"); err != nil {
		log.Fatalf("Failed to generate google token to work account: %v", err)
	}
	if err := google.GenerateToken("config/google-private-credential.json", "config/google-private-token.json"); err != nil {
		log.Fatalf("Failed to generate google token to private account: %v", err)
	}
}

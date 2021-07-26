package bot

import (
	"log"
	"net/http"
	"os"

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
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			token := r.URL.Query().Get("hub.verify_token")
			challenge := r.URL.Query().Get("hub.challenge")
			if os.Getenv("FB_VERIFY_TOKEN") == token {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(challenge))
			} else {
				w.WriteHeader(http.StatusForbidden)
			}
		case "POST":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unsupported method"))
		}
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

package bot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yuseinishiyama/stats/pkg/storage"
)

type bot struct {
	spreadsheet *storage.Spreadsheet
}

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
	i.spreadsheet = &storage.Spreadsheet{
		ID:         "1yG-Hzw4_U4wnEZMUToNGxb7v8_-Ab60BJrgTk6T4798",
		Credential: "config/google-private-credential.json",
		Token:      "config/google-private-token.json",
	}

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
			if err := i.handleMessage(w, r); err != nil {
				fmt.Println(err)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unsupported method"))
		}
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

func (i *bot) handleMessage(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	var message InputMessage
	err = json.Unmarshal(body, &message)
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	for _, entry := range message.Entry {
		event := entry.Messaging[0]
		if err = i.recordMood(event.Sender.ID, event.Message.Text); err != nil {
			err = fmt.Errorf("some messages were not handled properly")
		}
	}

	w.WriteHeader(200)
	return err
}

func (i *bot) sendMessage(rec string, message string) error {
	body := ResponseMessage{}
	body.Recipient.ID = rec
	body.Message.Text = message
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("https://graph.facebook.com/v2.6/me/messages?access_token=%s", os.Getenv("FB_ACCESS_TOKEN"))
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to post message. status code: %v", res.StatusCode)
	}

	return nil
}

func (i *bot) recordMood(rec string, score string) error {
	validationMsg := "Score must be integer between 1 and 5"
	parserdScore, err := strconv.ParseInt(score, 0, 0)
	if err != nil {
		if err = i.sendMessage(rec, validationMsg); err != nil {
			return err
		}
		return nil // recoverable error
	}

	if parserdScore < 1 || parserdScore > 5 {
		if err = i.sendMessage(rec, validationMsg); err != nil {
			return err
		}
		return nil // recoverable error
	}

	entry := storage.NewMoodEntry(int(parserdScore))
	if err = i.spreadsheet.Write(context.Background(), entry); err != nil {
		if err = i.sendMessage(rec, "Failed to record score on database"); err != nil {
			return err
		}
		return nil // recoverable error
	}

	if err = i.sendMessage(rec, "ACK"); err != nil {
		return err
	}

	return nil
}

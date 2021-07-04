package storage

import (
	"context"
	"log"

	"github.com/yuseinishiyama/stats/google"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Spreadsheet struct {}

func (s *Spreadsheet) Write(entry Entry) {
	ctx := context.Background()
	client := google.GetClient(ctx, "config/google-private-credential.json", "config/google-private-token.json")

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "1yG-Hzw4_U4wnEZMUToNGxb7v8_-Ab60BJrgTk6T4798"
	vr := sheets.ValueRange{}
	values := []interface{}{entry.Timestamp, entry.MailInboxWork, entry.MailInboxPrivate}
	vr.Values = append(vr.Values, values)
	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, "A1", &vr).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to append data to sheet: %v", err)
	}
}

package storage

import (
	"context"

	"github.com/yuseinishiyama/stats/pkg/google"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Spreadsheet struct {
	ID         string
	Credential string
	Token      string
}

func (s *Spreadsheet) Write(ctx context.Context, entry Entry) error {
	config, err := google.GetConfig(s.Credential)
	if err != nil {
		return err
	}
	client, err := google.GetClient(ctx, config, s.Token)
	if err != nil {
		return err
	}
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return err
	}
	vr := sheets.ValueRange{}
	values := []interface{}{entry.Timestamp, entry.Key, entry.Value}
	vr.Values = append(vr.Values, values)
	_, err = srv.Spreadsheets.Values.Append(s.ID, "A1", &vr).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	return err
}

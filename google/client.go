package google

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/sheets/v4"
)

func GetClient(ctx context.Context, config *oauth2.Config, token string) (*http.Client, error) {
	tok, err := tokenFromFile(token)
	if err != nil {
		return nil, err
	}
	return config.Client(ctx, tok), nil
}

func GetConfig(credential string) (*oauth2.Config, error) {
	b, err := ioutil.ReadFile(credential)
	if err != nil {
		return nil, err
	}
	// If modifying these scopes, delete your previously saved token.json.
	return google.ConfigFromJSON(b, gmail.GmailReadonlyScope, sheets.DriveScope)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

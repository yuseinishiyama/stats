package main

import (
	"github.com/yuseinishiyama/stats/provider"
	"github.com/yuseinishiyama/stats/storage"
)

func main() {
	spreadsheet := storage.Spreadsheet{}

	workGmail := &provider.Gmail{Credential: "config/google-work-credential.json", Token: "config/google-work-token.json"}
	mailInboxWork := storage.NewMailInboxWorkEntry(workGmail.Get())
	spreadsheet.Write(mailInboxWork)

	privateGmail := &provider.Gmail{Credential: "config/google-private-credential.json", Token: "config/google-private-token.json"}
	mailInboxPrivate := storage.NewMailInboxPrivateEntry(privateGmail.Get())
	spreadsheet.Write(mailInboxPrivate)
}

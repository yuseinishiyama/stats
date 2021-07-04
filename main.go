package main

import (
	"time"

	"github.com/yuseinishiyama/stats/provider"
	"github.com/yuseinishiyama/stats/storage"
)

func main() {
	workGmail := &provider.Gmail{Credential: "config/google-work-credential.json", Token: "config/google-work-token.json"}
	privateGmail := &provider.Gmail{Credential: "config/google-private-credential.json", Token: "config/google-private-token.json"}
	mailInboxWork := workGmail.Get()
	mailInboxPrivate := privateGmail.Get()
	entry := storage.Entry{Timestamp: time.Now().Unix(), MailInboxWork: mailInboxWork, MailInboxPrivate: mailInboxPrivate}
	spreadsheet := storage.Spreadsheet{}
	spreadsheet.Write(entry)
}

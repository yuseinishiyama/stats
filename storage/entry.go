package storage

import "time"

type Entry struct {
	Timestamp int64
	Key       string
	Value     int64
}

func NewMailInboxWorkEntry(value int64) Entry {
	return Entry{time.Now().Unix(), "mail_inbox_work", value}
}

func NewMailInboxPrivateEntry(value int64) Entry {
	return Entry{time.Now().Unix(), "mail_inbox_private", value}
}

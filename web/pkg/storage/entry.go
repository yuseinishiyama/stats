package storage

import "time"

type Entry struct {
	Timestamp int64
	Key       string
	Value     int
}

func NewMailInboxWorkEntry(value int) Entry {
	return Entry{time.Now().Unix(), "mail_inbox_work", value}
}

func NewMailInboxPrivateEntry(value int) Entry {
	return Entry{time.Now().Unix(), "mail_inbox_private", value}
}

func NewReadItLaterEntry(value int) Entry {
	return Entry{time.Now().Unix(), "read_it_later", value}
}

func NewChatSavedEntry(value int) Entry {
	return Entry{time.Now().Unix(), "chat_saved", value}
}

func NewMoodEntry(value int) Entry {
	return Entry{time.Now().Unix(), "mood", value}
}

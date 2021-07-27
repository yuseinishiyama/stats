package bot

type InputMessage struct {
	Object string
	Entry  []struct {
		ID        string
		Time      int64
		Messaging []struct {
			Postback struct {
				Title   string
				Payload string
			}
			Sender struct {
				ID string
			}
			Timestamp int64
			Message   struct {
				Mid        string
				Text       string
				QuickReply struct {
					Payload string
				}
			}
		}
	}
}

type ResponseMessage struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}

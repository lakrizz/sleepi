package models

import "time"

const (
	MSG_DOWNLOAD_STARTING = iota
	MSG_DOWNLOAD_FINISHED
	MSG_DOWNLOAD_ERROR
	MSG_SEARCH
	MSG_SEARCH_DONE
	MSG_EXIT
)

type Message struct {
	Status    int
	Payload   string
	Timestamp time.Time
}

func CreateMessage(payload string, status int) *Message {
	return &Message{Status: status, Payload: payload, Timestamp: time.Now()}
}

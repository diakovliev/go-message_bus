package message_bus

import (
    "github.com/google/uuid"
)

type MessageId struct {
    Id string `json:"id"`
}

func makeNewMessageIdString() string {
    return uuid.New().String()
}

func NewMessageId(messageIdString string) *MessageId {
    e := new(MessageId)
    e.Id = messageIdString
    return e
}

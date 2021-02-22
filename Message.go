package message_bus

import (
    "fmt"
    "encoding/json"
)

type Message struct {
    Id string           `json:"id"`
    UserId string       `json:"user_id"`
    Category string     `json:"category"`
    Payload interface{} `json:"payload"`
    Tag string          `json:"tag"`
}

func NewMessageExt(userId string, category string, payload interface{}, tag string) *Message {
    e := new(Message)
    e.Id        = makeNewMessageIdString()
    e.UserId    = userId
    e.Category  = category
    e.Payload   = payload
    e.Tag       = tag
    return e
}

func NewMessage(userId string, category string, payload interface{}) *Message {
    return NewMessageExt(userId, category, payload, "")
}

func (e Message) String() string {
    return fmt.Sprintf("{ Id: '%v', UserId: '%v', Category: '%v', tag: '%v' }", e.Id, e.UserId, e.Category, e.Tag)
}

func (e* Message) MarshalJson() ([]byte, error) {
    res, err := json.Marshal(e)
    return res, err
}

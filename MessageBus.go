package message_bus

import (
    "log"
    "fmt"
    "errors"

    "github.com/diakovliev/go-event_loop"
)

type MessageBus struct {
    Name string
    EventLoop *event_loop.EventLoop
    EventType uint64

    listeners *Listeners
}

func NewMessageBus(name string, eventLoop *event_loop.EventLoop, eventType uint64) *MessageBus {
    if eventLoop == nil {
        panic(errors.New("Can't create MessageBus without EventLoop instance!"))
    }

    mb := new(MessageBus)

    mb.Name         = name
    mb.EventLoop    = eventLoop
    mb.EventType    = eventType
    mb.listeners    = NewListeners()

    return mb.initHandler()
}

func (mb *MessageBus) initHandler() *MessageBus {
    mb.EventLoop.Subscribe(event_loop.NewEventsHandler(mb.EventType, mb.Name + "%handler", func(event event_loop.Event) bool {
        message, ok := event.Payload.(*Message)
        if !ok {
            panic(fmt.Errorf("Event payload is not a *Message!"))
        }

        mb.listeners.Range(func(key, value interface{}) bool {

            listener, ok := value.(*MessagesListener)
            if !ok {
                panic(fmt.Errorf("Value is not a *MessagesListener!"))
            }

            if !listener.matcher.Match(*message) {
                return true
            }

            listener.callback(*message)
            return true
        })

        return true
    }))
    return mb
}

func (mb *MessageBus) Subscribe(listener *MessagesListener) *MessagesListener {
    log.Printf("%s: Subscribe '%s' listener.", mb.Name, listener.userId)

    mb.listeners.Store(listener.id, listener)
    return listener
}

func (mb *MessageBus) Unsubscribe(listener *MessagesListener) *MessagesListener {
    log.Printf("%s: Unubscribe '%s' listener.", mb.Name, listener.userId)

    mb.listeners.Delete(listener.id)
    return listener
}

func (mb *MessageBus) Send(message *Message) {
    log.Printf("%s: Send message: %s", mb.Name, message.String())

    mb.EventLoop.Send(event_loop.NewEvent(mb.EventType, message))
}

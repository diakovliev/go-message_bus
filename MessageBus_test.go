package message_bus

import (
    "testing"

    "log"
    "github.com/diakovliev/go-event_loop"
)

const (
    TEST_MESSAGES = event_loop.EVENT_PRE_USER + 1
)

func TestCreateMessageBus(t *testing.T) {

    loop := event_loop.NewEventLoop("test_event_loop")

    mb := NewMessageBus("test_mb", loop, TEST_MESSAGES)
    if mb == nil {
        t.Fatal("Can't create MessageBus!")
    }

    mb.Subscribe(NewMessagesListener("listener", MatchAllMessages{}, func(message Message) bool {

        log.Printf("Recieved message: %s", message.String())

        return true
    }))

    mb.EventLoop.Start()

    mb.Send(NewMessage("test", "test_category", nil))
    mb.Send(NewMessage("test", "test_category", nil))
    mb.Send(NewMessage("test", "test_category", nil))
    mb.Send(NewMessage("test", "test_category", nil))
    mb.Send(NewMessage("test", "test_category", nil))

    mb.EventLoop.StopAndJoin(0)
}

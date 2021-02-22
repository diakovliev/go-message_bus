package message_bus

import "sync"

type Listeners struct {
    values []*MessagesListener
    guard sync.Mutex
}

func NewListeners() *Listeners {
    h := new(Listeners)
    h.values = make([]*MessagesListener, 0, 100)
    return h
}

func (h *Listeners) Store(name string, handler *MessagesListener) {
    h.guard.Lock()
    defer h.guard.Unlock()

    h.values = append(h.values, handler)
}

func (h *Listeners) Load(id string) (*MessagesListener, bool) {
    h.guard.Lock()
    defer h.guard.Unlock()

    for _, handler := range(h.values) {
        if handler.id == id {
            return handler, true
        }
    }

    return nil, false
}

func (h *Listeners) Delete(id string) {
    h.guard.Lock()
    defer h.guard.Unlock()

    idx := -1

    for iter_idx, handler := range(h.values) {
        if handler.id == id {
            idx = iter_idx
            break
        }
    }

    if idx != -1 {
        h.values = append(h.values[:idx], h.values[idx+1:]...)
    }
}

func (h *Listeners) Range(callback func (key, value interface{}) bool) {
    h.guard.Lock()
    defer h.guard.Unlock()

    for _, handler := range(h.values) {
        if !callback(handler.id, handler) {
            break
        }
    }
}

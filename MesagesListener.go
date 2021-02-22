package message_bus

type MessagesFilter interface {
    Match(Message) bool
}


// MessagesFilter impementation what will match all messages.
type MatchAllMessages struct {}

func (m MatchAllMessages) Match(Message) bool {
    return true
}


// MessagesFilter impementation what will match messages by exact UserId and Category.
type MatchMessage struct {
    UserId string
    Category string
    Tag string
}


func (m MatchMessage) Match(message Message) bool {
    matchedUserId   := len(m.UserId) == 0
    matchedCategory := len(m.Category) == 0
    matchedTag      := len(m.Tag) == 0

    if !matchedUserId {
        matchedUserId = m.UserId == message.UserId
    }
    if !matchedCategory {
        matchedCategory = m.Category == message.Category
    }
    if !matchedTag {
        matchedTag = m.Tag == message.Tag
    }

    return matchedUserId && matchedCategory && matchedTag
}


type MessagesListener struct {
    // Listener unique id. Will be generated automatically.
    id string

    // User id. Must be set by user code.
    userId string

    // Call 'callback' in separate goroutine.
    async bool

    // Messages matcher.
    matcher MessagesFilter

    // User callback.
    callback func(message Message) bool

    // Set of already recieved messages
    recievedMessages map[string]bool
}


func NewMessagesListener(userId string, matcher MessagesFilter, callback func(message Message) bool) *MessagesListener {
    listener := new(MessagesListener)
    listener.id = makeNewMessageIdString()
    listener.userId = userId
    listener.callback = callback
    listener.recievedMessages = make(map[string]bool)
    if matcher != nil {
        listener.matcher = matcher
    } else {
        listener.matcher = MatchAllMessages{}
    }
    listener.async = false
    return listener
}

func NewAsyncMessagesListener(userId string, matcher MessagesFilter, callback func(message Message) bool) *MessagesListener {
    listener := NewMessagesListener(userId, matcher, callback)
    listener.async = true
    return listener
}

func (l *MessagesListener) UserId() string {
    return l.userId
}

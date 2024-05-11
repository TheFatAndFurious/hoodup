package events

import "sync"

var (
	EventBroker = NewBroker()
)

type Broker struct {
	sync.Mutex
	clients map[chan string]struct{}
}


func NewBroker() *Broker {
	return &Broker{
		clients: make(map[chan string]struct{}),
	}
}

func(b *Broker) Subscribe() chan string {
	b.Lock()
	defer b.Unlock()

	ch := make(chan string, 10)
	b.clients[ch] = struct{}{}
	return ch
}

func(b *Broker) Unsubscribe(ch chan string) {
	b.Lock()
	defer b.Unlock()

	delete(b.clients, ch)
	close(ch)
}

func (b *Broker) Publish(msg string) {
    b.Lock()
    clients := make([]chan string, 0, len(b.clients))
    for ch := range b.clients {
        clients = append(clients, ch)
    }
    b.Unlock()

    for _, ch := range clients {
        select {
        case ch <- msg:
            // Message sent successfully
        default:
            // Log or handle the case where the message couldn't be sent
        }
    }
}

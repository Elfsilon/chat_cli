package broker

import "github.com/nats-io/nats.go"

type MessageBroker interface {
	Subscribe(subj string) (<-chan []byte, UnsubscribeFn, error)
	Publish(subj string, data []byte) error
}

type UnsubscribeFn func() error

type NatsMessageBroker struct {
	nc *nats.Conn
}

func NewNatsMessageBroker(nc *nats.Conn) *NatsMessageBroker {
	return &NatsMessageBroker{nc}
}

func (b *NatsMessageBroker) Subscribe(topic string) (<-chan []byte, UnsubscribeFn, error) {
	natsMsgs := make(chan *nats.Msg, 100)
	sub, err := b.nc.ChanSubscribe(topic, natsMsgs)
	if err != nil {
		return nil, nil, err
	}

	messages := make(chan []byte, 100)
	go func() {
		for m := range natsMsgs {
			messages <- m.Data
		}
	}()

	unsub := func() error {
		if err := sub.Unsubscribe(); err != nil {
			return err
		}
		close(natsMsgs)
		close(messages)

		return nil
	}

	return messages, unsub, nil
}

func (b *NatsMessageBroker) Publish(topic string, data []byte) error {
	return b.nc.Publish(topic, data)
}

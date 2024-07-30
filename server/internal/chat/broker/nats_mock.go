package broker

import "errors"

type MockMessageBroker struct {
	topicChans map[string]chan []byte
}

func NewMockMessageBroker() *MockMessageBroker {
	return &MockMessageBroker{
		topicChans: make(map[string]chan []byte, 0),
	}
}

func (b *MockMessageBroker) Subscribe(topic string) (<-chan []byte, UnsubscribeFn, error) {
	messages := make(chan []byte, 100)
	b.topicChans[topic] = messages

	unsub := func() error {
		close(b.topicChans[topic])
		delete(b.topicChans, topic)
		return nil
	}

	return messages, unsub, nil
}

func (b *MockMessageBroker) Publish(topic string, data []byte) error {
	ch, ok := b.topicChans[topic]
	if !ok {
		return errors.New("topic is not registered")
	}
	ch <- data
	return nil
}

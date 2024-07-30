package services

import (
	"context"
	"server/internal/chat/broker"
	"server/internal/chat/gen/chat"
	"server/internal/chat/repos"
	ctxutil "server/pkg/utils/context_utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type ServerStreamMock struct {
	grpc.ServerStream
	ctx context.Context
	que []*chat.ChatEvent
}

func (m *ServerStreamMock) Context() context.Context {
	return m.ctx
}

func (m *ServerStreamMock) Send(resp *chat.ChatEvent) error {
	m.que = append(m.que, resp)
	return nil
}

func initialize(t *testing.T) (context.Context, *ChatService, error) {
	t.Helper()

	r := repos.NewChatMockRepo()
	b := broker.NewMockMessageBroker()
	s := NewChatService(b, r)
	_, err := s.Init()

	ctx := context.WithValue(context.Background(), ctxutil.UserID, int64(1))

	return ctx, s, err
}

func TestCreateAndList(t *testing.T) {
	ctx, c, err := initialize(t)
	assert.NoError(t, err)

	_, err = c.Create(ctx, "normic", []int{0, 1, 2, 3})
	assert.NoError(t, err)

	list, err := c.List(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(list))
}

func TestDelete(t *testing.T) {
	ctx, c, err := initialize(t)
	assert.NoError(t, err)

	_, err = c.Create(ctx, "normic", []int{1, 2, 3})
	assert.NoError(t, err)

	list, err := c.List(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(list))

	err = c.Delete(ctx, 1)
	assert.NoError(t, err)

	list, err = c.List(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(list))
}

func TestConnectAndSend(t *testing.T) {
	ctx, service, err := initialize(t)

	ctx, cancel := context.WithCancel(ctx)
	assert.NoError(t, err)

	_, err = service.Create(ctx, "hz", []int{1, 2})
	assert.NoError(t, err)

	stream := &ServerStreamMock{ctx: ctx, que: make([]*chat.ChatEvent, 0)}

	go func() {
		time.Sleep(1 * time.Second)
		assert.Equal(t, 1, len(stream.que))

		err := service.SendMessage(ctx, 1, 1, "hz", "hz", 0)
		assert.NoError(t, err)

		time.Sleep(1 * time.Second)
		assert.Equal(t, 2, len(stream.que))

		cancel()
	}()

	err = service.Connect(stream.ctx, 1, 1, stream)
	assert.NoError(t, err)
}

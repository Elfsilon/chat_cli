package intc

import (
	"chat_cli/internal/app/services"
	"context"

	"google.golang.org/grpc"
)

type TokenProvider struct {
	auths *services.AuthService
}

func NewTokenProvider(s *services.AuthService) *TokenProvider {
	return &TokenProvider{auths: s}
}

func (t *TokenProvider) UnaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(t.auths.Context(ctx), method, req, reply, cc, opts...)
}

func (t *TokenProvider) StreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return streamer(t.auths.Context(ctx), desc, cc, method, opts...)
}

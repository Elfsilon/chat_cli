package grpcutils

import (
	"context"

	"google.golang.org/grpc"
)

type ServerStreamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func WrapServerStream(ctx context.Context, stream grpc.ServerStream) ServerStreamWrapper {
	return ServerStreamWrapper{
		ServerStream: stream,
		ctx:          ctx,
	}
}

func (s ServerStreamWrapper) Context() context.Context {
	return s.ctx
}

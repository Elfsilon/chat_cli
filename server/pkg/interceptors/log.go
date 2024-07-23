package intc

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func Log(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("method %q failed: %s", info.FullMethod, err)
	}
	return resp, err
}

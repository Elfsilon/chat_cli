package intc

import (
	"context"
	"errors"
	"fmt"
	"server/pkg/gen/auth"
	ctxutil "server/pkg/utils/context_utils"
	grpcutils "server/pkg/utils/grpc_utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	ErrMethodNotFound   = errors.New("method not found")
	ErrMetadataMissing  = errors.New("metadata is missing")
	ErrTokenMissing     = errors.New("token is missing")
	ErrInvalidStructure = errors.New("auth header structure is invalid")
	ErrNoPermission     = errors.New("no permission")
)

type AuthGuard struct {
	client auth.AuthClient
}

func NewAuthGuard(client auth.AuthClient) *AuthGuard {
	return &AuthGuard{client}
}

func (a *AuthGuard) UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	ctx, err := a.provideAuthContext(ctx, info.FullMethod)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func (a *AuthGuard) StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx, err := a.provideAuthContext(stream.Context(), info.FullMethod)
	if err != nil {
		return err
	}
	updatedStream := grpcutils.WrapServerStream(ctx, stream)
	return handler(srv, updatedStream)
}

func (a *AuthGuard) provideAuthContext(ctx context.Context, method string) (context.Context, error) {
	requestData := &auth.CheckResourceRequest{FullMethod: method}

	md, _ := metadata.FromIncomingContext(ctx)
	outCtx := metadata.NewOutgoingContext(ctx, md)

	res, err := a.client.CheckResource(outCtx, requestData)
	if err != nil {
		return ctx, err
	}
	if !res.HasAccess {
		return ctx, fmt.Errorf(res.Reason)
	}

	return context.WithValue(ctx, ctxutil.UserID, res.GetUserID()), nil
}

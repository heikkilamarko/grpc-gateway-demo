package internal

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewApiKeyUnaryServerInterceptor(name string, apiKey string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			slog.Info("incoming metadata not found")
			return nil, errors.New("incoming metadata not found")
		}

		apiKeys := md[name]
		if len(apiKeys) == 0 {
			slog.Info("api key not found")
			return nil, errors.New("api key not found")
		}

		if apiKeys[0] != apiKey {
			slog.Info("invalid api key")
			return nil, errors.New("invalid api key")
		}

		return handler(ctx, req)
	}
}

func NewApiKeyMetadataAnnotator(name string, apiKey string) func(context.Context, *http.Request) metadata.MD {
	return func(context.Context, *http.Request) metadata.MD {
		md := metadata.MD{}
		md.Set(name, apiKey)
		return md
	}
}

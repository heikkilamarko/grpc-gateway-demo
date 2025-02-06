package main

import (
	"calculator-demo/internal"
	"calculator-demo/internal/calculator/v1"
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		serverAddress            = os.Getenv("SERVER_ADDRESS")
		calculatorServiceAddress = os.Getenv("CALCULATOR_SERVICE_ADDRESS")
		ctx                      = context.Background()
		mux                      = runtime.NewServeMux()
		opts                     = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	)

	if err := calculator.RegisterCalculatorServiceHandlerFromEndpoint(ctx, mux, calculatorServiceAddress, opts); err != nil {
		slog.Error("register calculator service handler", "error", err)
		os.Exit(1)
	}

	server := &http.Server{
		Addr:    serverAddress,
		Handler: internal.RequestLogger(mux),
	}

	slog.Info("listen", "addr", serverAddress)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error("listen", "error", err)
		os.Exit(1)
	}
}

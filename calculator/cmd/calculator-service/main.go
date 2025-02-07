package main

import (
	"calculator-demo/internal"
	"calculator-demo/internal/calculator/v1"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	var (
		serverAddress = os.Getenv("SERVER_ADDRESS")
		apiKeyName    = os.Getenv("API_KEY_NAME")
		apiKey        = os.Getenv("API_KEY")
	)

	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		slog.Error("create listener", "error", err)
		os.Exit(1)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(internal.NewApiKeyUnaryServerInterceptor(apiKeyName, apiKey)))

	calculator.RegisterCalculatorServiceServer(server, &internal.CalculatorServiceServer{})

	slog.Info("listen", "addr", serverAddress)

	if err := server.Serve(listener); err != nil {
		slog.Error("listen", "error", err)
		os.Exit(1)
	}
}

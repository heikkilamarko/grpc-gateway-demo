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
	addr := os.Getenv("SERVER_ADDRESS")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("create listener", "error", err)
		os.Exit(1)
	}

	server := grpc.NewServer()

	calculator.RegisterCalculatorServiceServer(server, &internal.CalculatorServiceServer{})

	slog.Info("listen", "addr", addr)

	if err := server.Serve(listener); err != nil {
		slog.Error("listen", "error", err)
		os.Exit(1)
	}
}

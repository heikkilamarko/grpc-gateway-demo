package main

import (
	"calculator-demo/internal/calculator"
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(os.Getenv("CALCULATOR_SERVICE_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("connect to calculator service", "error", err)
		os.Exit(1)
	}

	gwmux := runtime.NewServeMux()

	err = calculator.RegisterCalculatorHandler(context.Background(), gwmux, conn)
	if err != nil {
		slog.Error("register calculator handler", "error", err)
		os.Exit(1)
	}

	addr := os.Getenv("SERVER_ADDRESS")

	server := &http.Server{
		Addr:    addr,
		Handler: requestLogger(gwmux),
	}

	slog.Info("listen", "addr", addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error("listen", "error", err)
		os.Exit(1)
	}
}

func requestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		slog.Info("request", "path", r.URL.Path, "hostname", hostname)
		h.ServeHTTP(w, r)
	})
}

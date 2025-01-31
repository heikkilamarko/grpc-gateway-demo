package main

import (
	"calculator-demo/proto/calculator"
	"context"
	"log"
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
		log.Fatalln("dial server", err)
	}

	gwmux := runtime.NewServeMux()

	err = calculator.RegisterCalculatorHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("register gateway", err)
	}

	addr := os.Getenv("SERVER_ADDRESS")

	gwServer := &http.Server{
		Addr:    addr,
		Handler: requestLogger(gwmux),
	}

	log.Println("serving grpc gateway on " + addr)
	log.Fatalln(gwServer.ListenAndServe())
}

func requestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		slog.Info("request", "path", r.URL.Path, "hostname", hostname)
		h.ServeHTTP(w, r)
	})
}

package main

import (
	"calculator-demo/proto/calculator"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	addr := os.Getenv("SERVER_ADDRESS")

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("listen:", err)
	}

	s := grpc.NewServer()

	calculator.RegisterCalculatorServer(s, &server{})

	log.Println("serving grpc on " + addr)
	log.Fatalln(s.Serve(l))
}

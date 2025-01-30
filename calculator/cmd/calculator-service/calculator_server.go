package main

import (
	"calculator-demo/proto/calculator"
	"context"
)

type server struct {
	calculator.UnimplementedCalculatorServer
}

func (s *server) Add(_ context.Context, req *calculator.AddRequest) (*calculator.AddReply, error) {
	result := req.A + req.B
	return &calculator.AddReply{Result: result}, nil
}

func (s *server) Subtract(_ context.Context, req *calculator.SubtractRequest) (*calculator.SubtractReply, error) {
	result := req.A - req.B
	return &calculator.SubtractReply{Result: result}, nil
}

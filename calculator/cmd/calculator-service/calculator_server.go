package main

import (
	"calculator-demo/internal/calculator"
	"context"
)

type calculatorServer struct {
	calculator.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(_ context.Context, req *calculator.AddRequest) (*calculator.AddReply, error) {
	result := req.A + req.B
	return &calculator.AddReply{Result: result}, nil
}

func (s *calculatorServer) Subtract(_ context.Context, req *calculator.SubtractRequest) (*calculator.SubtractReply, error) {
	result := req.A - req.B
	return &calculator.SubtractReply{Result: result}, nil
}

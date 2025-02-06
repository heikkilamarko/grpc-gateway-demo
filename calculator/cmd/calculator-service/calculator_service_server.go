package main

import (
	"calculator-demo/internal/calculator/v1"
	"context"
)

type calculatorServiceServer struct {
	calculator.UnimplementedCalculatorServiceServer
}

func (s *calculatorServiceServer) Add(_ context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	result := req.A + req.B
	return &calculator.AddResponse{Result: result}, nil
}

func (s *calculatorServiceServer) Subtract(_ context.Context, req *calculator.SubtractRequest) (*calculator.SubtractResponse, error) {
	result := req.A - req.B
	return &calculator.SubtractResponse{Result: result}, nil
}

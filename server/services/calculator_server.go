package services

import (
	context "context"
	"fmt"
)

type calculatorServer struct {
	UnimplementedCalculatorServer
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	result := fmt.Sprintf("Hello %v", req.Name)
	res := HelloResponse{
		Result: result,
	}

	return &res, nil
}

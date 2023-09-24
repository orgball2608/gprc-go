package main

import (
	"context"
	"fmt"
	"learn/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	log.Println("sum called...")
	resp := &calculatorpb.AddResponse{
		Result: req.GetNum2() + req.GetNum1(),
	}

	return resp, nil
}

func (*server) Subtract(ctx context.Context, req *calculatorpb.SubtractRequest) (*calculatorpb.SubtractResponse, error) {
	log.Println("substract called...")
	substract := &calculatorpb.SubtractResponse{
		Result: req.GetNum1() - req.GetNum2(),
	}

	return substract, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4567")

	if err != nil {
		log.Fatalf("Failed to listen on port 4567: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("calculator is running... in port 4567")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"learn/calculator/calculatorpb"
	"log"
	"net"
	"time"

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

func (*server) PrimeNumberDecomposition(req *calculatorpb.PNDRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Println("PrimeNumberDecomposition called...")
	k := int32(2)
	N := req.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N / k
			stream.Send(&calculatorpb.PNDResponse{
				Result: k,
			})
			time.Sleep(1000 * time.Millisecond)
		} else {
			k++
			log.Printf("k increase to %v", k)
		}
	}
	return nil
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

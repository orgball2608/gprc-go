package main

import (
	"context"
	"fmt"
	"io"
	"learn/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:4567", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)

	// 	callAdd(client)
	//
	// 	callSubstract(client)

	callPND(client)
}

func callAdd(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("calling add...")
	req := &calculatorpb.AddRequest{
		Num1: 10,
		Num2: 5,
	}

	resp, err := c.Add(context.Background(), req)

	if err != nil {
		log.Fatalf("err while calling add %v", err)
	}

	log.Printf("response from add %v", resp)
}

func callSubstract(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("calling substract...")
	req := &calculatorpb.SubtractRequest{
		Num1: 10,
		Num2: 5,
	}

	resp, err := c.Subtract(context.Background(), req)

	if err != nil {
		log.Fatalf("err while calling substract %v", err)
	}

	log.Printf("response from substract %v", resp)
}

func callPND(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("calling PrimeNumberDecomposition...")
	stream, err := c.PrimeNumberDecomposition(context.Background(), &calculatorpb.PNDRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("err while calling PrimeNumberDecomposition %v", err)
	}

	for {
		resp, recvErr := stream.Recv()
		if recvErr == io.EOF {
			log.Println("server finish streaming")
			return
		}

		if recvErr != nil {
			log.Fatalf("callPND recvErr %v", recvErr)
		}

		log.Printf("prime number %v", resp.GetResult())
	}
}

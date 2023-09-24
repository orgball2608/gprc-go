package main

import (
	"context"
	"fmt"
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

	callAdd(client)

	callSubstract(client)
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

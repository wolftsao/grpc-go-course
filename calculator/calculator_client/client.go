package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/wolftsao/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Calculator Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)

	// doServerStreaming(c)

	doClientStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}

		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a ComputeAverage Client Streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	numbers := []int32{2, 3, 5, 9, 54, 23}

	for _, number := range numbers {
		fmt.Printf("Sending number: %d\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{Number: number})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	fmt.Printf("The Average is: %v", res.GetAverage())
}

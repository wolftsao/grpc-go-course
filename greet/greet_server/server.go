package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wolftsao/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

var server greetpb.UnimplementedGreetServiceServer

func main() {
	fmt.Println("Hellow world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

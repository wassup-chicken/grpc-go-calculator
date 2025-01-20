package main

import (
	"context"
	"io"
	"log"

	pb "github.com/wassup-chicken/grpc-exercise-1/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{Input: 12390392840}

	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Prime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", msg.Result)
	}
}

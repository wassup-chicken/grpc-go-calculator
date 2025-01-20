package main

import (
	"context"
	"log"
	"time"

	pb "github.com/wassup-chicken/grpc-exercise-1/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	numbers := []int64{3, 5, 9, 54, 23}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)
		stream.Send(&pb.AverageRequest{
			Input: number,
		})

		time.Sleep(1 * time.Second)

	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)

}

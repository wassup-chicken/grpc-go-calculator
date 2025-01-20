package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/wassup-chicken/grpc-exercise-1/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	numbers := []int64{4, 7, 2, 19, 4, 6, 32}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Input: number,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving %v\n", err)
				break
			}

			log.Printf("Received a new maximum: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}

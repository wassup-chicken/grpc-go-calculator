package main

import (
	"io"
	"log"

	pb "github.com/wassup-chicken/grpc-exercise-1/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("Average function was invoked with: %v\n", stream)
	count := int64(0)
	sum := int64(0)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{Result: float64(sum) / float64(count)})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving number: %d\n", req.Input)
		sum += int64(req.Input)
		count++
	}
}

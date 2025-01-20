package main

import (
	"io"
	"log"

	pb "github.com/wassup-chicken/grpc-exercise-1/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")
	var maximum int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if number := req.Input; number > maximum {
			maximum = number
			err = stream.Send(&pb.MaxResponse{
				Result: maximum,
			})
		}

		if err != nil {
			log.Fatalf("Error while sending data to cleint: %v\n", err)
		}

	}
}

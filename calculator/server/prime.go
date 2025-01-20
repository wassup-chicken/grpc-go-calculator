package main

import (
	"log"

	pb "github.com/wassup-chicken/grpc-exercise-1/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Prime function was invoked with: %v\n", in)

	divisor := int64(2)
	number := in.Input

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number = number / divisor
		} else {
			divisor++
		}

	}

	return nil

}

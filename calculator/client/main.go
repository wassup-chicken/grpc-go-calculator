package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/wassup-chicken/grpc-exercise-1/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//Unary
	// doSum(c)

	//Stream from Servere
	// doPrime(c)

	//Stream from Client
	// doAverage(c)

	//Bidirectional
	// doMax(c)

	//Error Handling
	doSqrt(c, -10)
}

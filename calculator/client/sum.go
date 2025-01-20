package main

import (
	"context"
	"log"

	pb "github.com/wassup-chicken/grpc-exercise-1/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	firstNum := int32(12)
	secondNum := int32(15)

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNum:  firstNum,
		SecondNum: secondNum,
	})

	if err != nil {
		log.Fatalf("Could not Sum: %v\n", err)
	}

	log.Printf("The Sum of %v and %v is %v\n", firstNum, secondNum, res.Result)
}

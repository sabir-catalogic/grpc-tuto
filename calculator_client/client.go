package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"test_go/calculator/calculatorpb"
)

func main() {
	fmt.Println("Client started....")
	cc,err:=grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("could not connect: %v",err)
	}
	defer cc.Close()
	c:=calculatorpb.NewCalculatorServiceClient(cc)

	//doUnary(c)
	doServerStreaming(c)
}

//func doUnary(c calculatorpb.CalculatorServiceClient)  {
//	fmt.Println("Starting DoUnary rpc...")
//	req:=&calculatorpb.SumRequest{
//		FirstNumber: 5,
//		SecondNumber: 6,
//	}
//	res,err:=c.Sum(context.Background(),req)
//	if err !=nil{
//		log.Fatalf("error while calling sum rpc:%v",err)
//	}
//	log.Printf("Response from Sum : %v",res.SumResult)
//}


func doServerStreaming(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("Starting PrimeNumberDecomposition ServerStreaming rpc...")
	req:=&calculatorpb.PrimeNumberDecompositonRequest{
		Number: 26,
	}

	stream,err:=c.PrimeNumberDecomposition(context.Background(),req)
	if err !=nil{
		log.Fatalf("error while calling PrimeDecomposition rpc:%v",err)
	}
	for {
		res,err:=stream.Recv()
		if err==io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("something went wrong %v",err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
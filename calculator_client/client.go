package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	//doServerStreaming(c)
	doClientStreaming(c)
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

//
//func doServerStreaming(c calculatorpb.CalculatorServiceClient)  {
//	fmt.Println("Starting PrimeNumberDecomposition ServerStreaming rpc...")
//	req:=&calculatorpb.PrimeNumberDecompositonRequest{
//		Number: 26,
//	}
//
//	stream,err:=c.PrimeNumberDecomposition(context.Background(),req)
//	if err !=nil{
//		log.Fatalf("error while calling PrimeDecomposition rpc:%v",err)
//	}
//	for {
//		res,err:=stream.Recv()
//		if err==io.EOF{
//			break
//		}
//		if err!=nil{
//			log.Fatalf("something went wrong %v",err)
//		}
//		fmt.Println(res.GetPrimeFactor())
//	}
//}

func doClientStreaming(c calculatorpb.CalculatorServiceClient){
	fmt.Printf("Starting to do compute client streaming rpc ...")

	stream,err:=c.ComputeAverage(context.Background())
	if err!=nil{
		log.Fatalf("error while sending stream :%v",err)
	}
	numbers:=[]int32{2,4,6,3,8}

	for _,number :=range numbers{
		fmt.Printf("sending number:%v \n",number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}
	res,err:=stream.CloseAndRecv()
	if err!=nil{
		log.Fatalf("error while receiving the response :%v",err)
	}
	fmt.Printf("Average is %v",res.GetAverage())

}
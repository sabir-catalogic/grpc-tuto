package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"test_go/calculator/calculatorpb"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error){
	fmt.Printf("Recevied sum rpc: %v \n",req)
	firstNumber:=req.FirstNumber
	secondNumber:=req.SecondNumber
	sum:=firstNumber+secondNumber
	res:=&calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res,nil
}

func main() {
	fmt.Println("Server Started")
	lis,err:=net.Listen("tcp","0.0.0.0:50051")
	if err !=nil{
		log.Fatalf("Server Failed to Listen %v",err)
	}
	s:=grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s,&server{})
	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("failed to serve:%v",err)
	}
}

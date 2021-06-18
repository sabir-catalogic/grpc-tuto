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

func (*server) 	PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositonRequest,stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error{
	fmt.Printf("Recevied PrimeNumberDecomposition rpc: %v \n",req)
	number:= req.GetNumber()
	divisor:=int64(2)

	for number>1{
		if number%divisor==0{
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number=number/divisor
		} else{
			divisor++
			fmt.Printf("Divisor= %v \n",divisor)
		}
	}
	return nil
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

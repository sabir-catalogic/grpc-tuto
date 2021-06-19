package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
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

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error  {
	fmt.Printf("REceived ComputeAverage rpc \n")
	sum:=float64(0)
	count:=0

	for{
		req,err:=stream.Recv()
		if err==io.EOF{
			average:=float64(sum)/float64(count)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: average,
			})
		}
		if err!=nil{
			log.Fatalf("Error with reading client :%v",err)
		}
		sum+=float64(req.GetNumber())
		count++
	}
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

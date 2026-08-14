package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpcstreamclient/proto/gen"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := mainpb.NewCalculatorClient(conn)

	ctx := context.Background()
	req := &mainpb.FibonacciRequest{
		N: 10,
	}

	stream, err := client.GenerateFibonacci(ctx,req)
	if err != nil {
		log.Fatalln("Error calling GenerateFibonacci", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			break
		}
		if err != nil {
			log.Fatalln("Error reciving data from GenerateFibonacci", err)
		}

		log.Println("Fibonacci number: ",resp.GetNumber())
	}
}

package main

import (
	"context"
	"io"
	"log"
	"time"

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

	// -------------- SERVER SIDE STREAMING ENDS
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
	// -------------- SERVER SIDE STREAMING ENDS

	// -------------- CLIENT SIDE STREAMING STARTS
	stream1, err := client.SendNumbers(ctx)
	if err != nil {
		log.Fatalln("Error creating stream", err)
	}

	for num := range 9 {
		err := stream1.Send(&mainpb.NumberRequest{Number: int32(num)})
		if err != nil {
			log.Fatalln("Error sending number", err)
		}

		time.Sleep(time.Second)
	}

	res, err := stream1.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error reciving response", err)
	}

	log.Println("SUM:", res.Sum)
	// -------------- CLIENT SIDE STREAMING ENDS
}

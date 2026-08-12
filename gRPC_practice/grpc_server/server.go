package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"
	farewellpb "simplegrpcserver/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	// _ "google.golang.org/grpc/encoding/gzip"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedBidFarewellServer
}

type serverGreeter struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	// md, ok := metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	log.Println("no metadata received")
	// }
	// log.Println("Metadata:", md)
	// val, ok := md["authorization"]
	// if !ok {
	// 	log.Println("No value with auth key in metadata")
	// }
	// log.Println("Authorization:", val[0])

	// // Set response headers
	// responseHeaders := metadata.Pairs("test", "testvalue", "test2", "testing2")
	// err := grpc.SendHeader(ctx, responseHeaders)
	// if err != nil {
	// 	return nil, err
	// }
	sum := req.A + req.B
	log.Println("Sum:", sum)

	// trailer := metadata.Pairs("testTrailer", "testtrailerVal", "testTrailer2", "testtrailerVal2")
	// grpc.SetTrailer(ctx, trailer)
	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

func (s *serverGreeter) Add(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s. Nice to receive request from you", req.Name),
	}, nil
}

func (s *serverGreeter) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s. Nice to receive request from you", req.Name),
	}, nil
}

func (s *server) BidGoodBye(ctx context.Context, req *farewellpb.GoodByeRequest) (*farewellpb.GoodByeResponse, error) {
	return &farewellpb.GoodByeResponse{
		Message: fmt.Sprintf("Goodbye %s!. Nice to have received request from you. Farewell my friend!", req.Name),
	}, nil
}

func main() {

	cert := "cert.pem"
	key := "key.pem"

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalln("Failed to load credentials", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculateServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &serverGreeter{})
	pb.RegisterBidFarewellServer(grpcServer, &server{})
	// farewellpb.RegisterAufWiedersehenServer(grpcServer, &server{})

	reflection.Register(grpcServer)
	log.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}

}

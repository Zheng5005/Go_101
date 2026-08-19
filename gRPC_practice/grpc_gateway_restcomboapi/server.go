package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"

	pb "grpcgatewayproject/proto/gen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// server is used to implement mainapi.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// Greet implements mainapi.GreeterServer
func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Validate the incoming request
	err := req.Validate()
	if err != nil {
		log.Printf("Validation failed: %v", err)
		// Return an invalid argument error if validation fails
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request: %v", err)
	}

	return &pb.HelloResponse{Message: "Hello, " + req.GetName()}, nil
}

func runGRPCSever(certFile, keyFile string) {
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatal("Failed to load TLS certificates:", err)
	}

	// Create a listener on TCP port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	log.Println("gRPC Server is running on port :50051...")
	// Start the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func loadTLSCredentials(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal("Failed to load certificates:", err)
	}
	return cert
}

// REST API
func runGatewayServer(certFile, keyFile string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	}))}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatal("Failed to register gRPC-Gateway handler:", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{loadTLSCredentials(certFile, keyFile)},
	}

	server := &http.Server{
		Addr:      ":8080",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	log.Println("HTTPS Server is running on port: 8080...")
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal("Failed to start HTTP Server:", err)
	}

	// // -------- Without TLS
	// log.Println("HTTP Server is running on port: 8080...")
	// err = http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	log.Fatal("Failed to serve HTTP:", err)
	// }
}

func main() {
	cert := "cert.pem"
	key := "key.pem"

	go runGRPCSever(cert, key)
	runGatewayServer(cert, key)
}

package grpcserver

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/andrew-j-price/journey/messages"
	"google.golang.org/grpc"
)

const port = ":50000"

type server struct{}

func (s *server) SayHello(ctx context.Context, req *messages.HelloRequest) (*messages.HelloResponse, error) {
	fmt.Printf("Received name: %s\n", req.Name)
	return &messages.HelloResponse{Message: "Hello " + req.Name}, nil
}

func Main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	messages.RegisterHelloServiceServer(s, &server{})
	fmt.Println("Starting...")

	s.Serve(lis)
}

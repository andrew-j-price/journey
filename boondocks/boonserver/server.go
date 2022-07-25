package boonserver

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/andrew-j-price/journey/boondocks/messages"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("boondocks-port", 4444, "gRPC Server listening port")
)

type server struct {
	pb.UnimplementedBoonServiceServer
}

func (s *server) PerformHelloWorld(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func Main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failure listening on: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBoonServiceServer(s, &server{})
	log.Printf("Listening on: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failure serving: %v", err)
	}
}

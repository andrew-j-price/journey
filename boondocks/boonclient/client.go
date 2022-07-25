package boonclient

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/andrew-j-price/journey/boondocks/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "woods"
)

var (
	addr = flag.String("boondocks-addr", "localhost:4444", "gRPC address to connect to")
	name = flag.String("boondocks-name", defaultName, "gRPC Name to use")
)

func Main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failure connecting to: %v", err)
	}
	defer conn.Close()
	c := pb.NewBoonServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PerformHelloWorld(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("Error on PerformHelloWorld: %v", err)
	}
	log.Printf("Hi there: %s", r.GetMessage())
}

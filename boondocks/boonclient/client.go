package boonclient

import (
	"context"
	"flag"
	"time"

	pb "github.com/andrew-j-price/journey/boondocks/messages"
	"github.com/andrew-j-price/journey/helpers"
	"github.com/andrew-j-price/journey/logger"
	"github.com/andrew-j-price/journey/rps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "woods"
)

var (
	addr = flag.String("boondocks-addr", "localhost:4444", "gRPC address to connect to")
	name = flag.String("boondocks-name", defaultName, "Hello World name to use")
)

func init() {
	flag.Bool("boondocks-rps", false, "Play a game of rock-paper-scissors")
}

func Main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal.Fatalf("Failure connecting to: %v", err)
	}
	defer conn.Close()
	c := pb.NewBoonServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// gRPC - Hello World function
	if helpers.IsFlagPassed("boondocks-name") {
		r, err := c.PerformHelloWorld(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			logger.Fatal.Fatalf("Error on PerformHelloWorld: %v", err)
		}
		logger.Info.Printf("HelloWorld Result: %s", r.GetMessage())
	}

	// gRPC - Play Rock, Paper, Scissors
	if helpers.IsFlagPassed("boondocks-rps") {
		/* // TODO: evaluate enum from protobuf
		throw := "SCISSORS"
		value, ok := pb.RpsChoice2_Play_value[throw]
		if !ok {
			panic("invalid enum value")
		}
		fmt.Printf("value: is of type: %v, with value: %v\n", reflect.TypeOf(value), value)
		*/
		toss := rps.RandomChoice()
		logger.Info.Printf("Tossing: %s", toss)
		rps, err := c.PlayRps(ctx, &pb.RpsChoice{Throw: toss})
		if err != nil {
			logger.Fatal.Fatalf("Error on PlayRps: %v", err)
		}
		// NOTE: use either output method below
		// logger.Info.Printf("Results: %s", rps)
		logger.Info.Printf("Results: %s", helpers.PrettyPrintStruct(rps))
	}
}

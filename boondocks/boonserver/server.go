package boonserver

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "github.com/andrew-j-price/journey/boondocks/messages"
	"github.com/andrew-j-price/journey/logger"
	"github.com/andrew-j-price/journey/rps"
	"google.golang.org/grpc"
)

var (
	port      = flag.Int("boondocks-port", 4444, "gRPC Server listening port")
	gameScore = new(rps.TheScore)
)

type server struct {
	pb.UnimplementedBoonServiceServer
}

func (s *server) PerformHelloWorld(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	logger.Info.Printf("PerformHelloWorld - received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s *server) PlayRps(ctx context.Context, client *pb.RpsChoice) (*pb.RpsScore, error) {
	logger.Info.Printf("PlayRps - received: %v", client.Throw)
	rps.PlayGame(gameScore, client.Throw, rps.RandomChoice())
	return &pb.RpsScore{UserWins: strconv.Itoa(gameScore.UserWins), CompWins: strconv.Itoa(gameScore.CompWins), Draws: strconv.Itoa(gameScore.Draws)}, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failure listening on: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBoonServiceServer(s, &server{})
	logger.Info.Printf("Listening on: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failure serving: %v", err)
	}
}

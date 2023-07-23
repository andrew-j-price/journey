package main

import (
	"flag"
	"os"

	"github.com/andrew-j-price/journey/pkg/boondocks/boonclient"
	"github.com/andrew-j-price/journey/pkg/boondocks/boonserver"
	"github.com/andrew-j-price/journey/pkg/logger"
)

func init() {
	logger.Logger()
}

func main() {
	runGrpcClient := flag.Bool("client", false, "start GRPC client")
	runGrpcServer := flag.Bool("server", false, "start GRPC server")
	flag.Parse()

	if !*runGrpcClient && !*runGrpcServer {
		flag.Usage()
		logger.Error.Println("\n\nERROR: Must specify either -client or -server flag")
		os.Exit(1)
	}

	if *runGrpcClient {
		boonclient.Main()
		os.Exit(0)
	}
	if *runGrpcServer {
		boonserver.Main()
		os.Exit(3)
	}
}

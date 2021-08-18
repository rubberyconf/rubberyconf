package main

import (
	"fmt"

	"log"
	"net"

	"github.com/rubberyconf/rubberyconf/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/grpcapi/servers"
	"google.golang.org/grpc"
)

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	grpcapipb.RegisterRubberyConfServiceServer(s, &servers.ConfServer{})
	grpcapipb.RegisterRubberyFeatureServiceServer(s, &servers.FeatureServer{})

	s.Serve(lis)
}

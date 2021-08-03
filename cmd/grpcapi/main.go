package main

import (
	"context"
	"fmt"

	"log"
	"net"

	"github.com/rubberyconf/rubberyconf/grpcapi/grpcapipb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) get(ctx context.Context, request *grpcapipb.RubberyConfRequest) (*grpcapipb.RubberyConfResponse, error) {
	/*name := request.Name
	response := &grpcapipb.rubberyConfResponse{
		Greeting: "Hello " + name,
	}
	return response, nil*/
	return nil, nil
}
func (*server) create(ctx context.Context, request *grpcapipb.RubberyConfRequest) (*grpcapipb.RubberyConfResponse, error) {
	/*name := request.Name
	response := &hellopb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil*/
	return nil, nil
}
func (*server) delete(ctx context.Context, request *grpcapipb.RubberyConfRequest) (*grpcapipb.RubberyConfResponse, error) {
	/*name := request.Name
	response := &hellopb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil*/
	return nil, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	grpcapipb.RubberyConfServiceServer
	grpcapipb.RegisterRubberyConfServiceServer(s, &server{})

	s.Serve(lis)
}

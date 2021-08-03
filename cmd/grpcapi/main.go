package main

import (
	"context"
	"fmt"

	"log"
	"net"

	"github.com/rubberyconf/rubberyconf/internal/business"

	"github.com/rubberyconf/rubberyconf/grpcapi/grpcapipb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureFullResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}
	result, content, typeContent := logic.GetFeature(vars)

	response := &grpcapipb.FeatureFullResponse{
		Status: result,
		Value:  content,
	}
	return response, nil
}
func (*server) create(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureResponse, error) {
	var logic business.Business
	name := request.Name

	vars := map[string]string{"feature": name}
	result, content, typeContent := logic.CreateFeature(vars)
	response := &grpcapipb.FeatureResponse{
		Status: result,
	}
	return response, nil
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

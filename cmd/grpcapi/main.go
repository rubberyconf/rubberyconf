package main

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"net"

	"github.com/rubberyconf/rubberyconf/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/internal/business"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureFullResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}
	result, content, _ := logic.GetFeature(vars)

	bytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	serialized := string(bytes)
	response := &grpcapipb.FeatureFullResponse{
		Status: grpcapipb.StatusType(result),
		Value:  serialized,
	}
	return response, nil
}

func mapper(in *grpcapipb.FeatureCreationRequestDefaultCls) feature.FeatureDefinition {

	var result = feature.FeatureDefinition{}
	result.Default.Value.Data = in.Value.Data // to be reviewed
	result.Default.Value.Type = in.Value.Type
	//TODO: complete other fields

	return result

}

func (*server) Create(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureResponse, error) {
	var logic business.Business
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapper(request.Default)

	result, _ := logic.CreateFeature(vars, ruberConf)

	response := &grpcapipb.FeatureResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, nil
}
func (*server) Delete(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, _ := logic.DeleteFeature(vars)

	response := &grpcapipb.FeatureResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	grpcapipb.RegisterRubberyConfServiceServer(s, &server{})

	s.Serve(lis)
}

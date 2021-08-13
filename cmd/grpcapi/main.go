package main

import (
	"context"
	"fmt"

	"log"
	"net"

	"github.com/rubberyconf/rubberyconf/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/internal/business"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"google.golang.org/grpc"
)

type ConfServer struct {
}
type FeatureServer struct {
}

func (*ConfServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureFullResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}
	result, content := logic.GetFeatureFull(vars)

	/*bytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}*/

	//serialized := string(bytes)
	// TODO: mapping fields one by one
	response := &grpcapipb.FeatureFullResponse{
		Status:     grpcapipb.StatusType(result),
		Defaultttl: content.Default.TTL,
		//DefaultValue: content.Default.Value.Data
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

func (*ConfServer) Create(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureResponse, error) {
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
func (*ConfServer) Delete(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, _ := logic.DeleteFeature(vars)

	response := &grpcapipb.FeatureResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, nil
}

func (*FeatureServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureShortResponse, error) {
	// TODO: implement it
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
	grpcapipb.RegisterRubberyConfServiceServer(s, &ConfServer{})
	grpcapipb.RegisterRubberyFeatureServiceServer(s, &FeatureServer{})

	s.Serve(lis)
}

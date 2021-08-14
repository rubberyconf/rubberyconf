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

//TODO: complete full mapping
func (*ConfServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureFullResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}
	result, content, err := logic.GetFeatureFull(vars)

	response := &grpcapipb.FeatureFullResponse{
		Status:     grpcapipb.StatusType(result),
		FeatureDef: mapperFromFeatureDef(content),
	}
	return response, err
}

func mapperToFeatureDef(in *grpcapipb.FeatureCreationRequestDefaultCls) feature.FeatureDefinition {

	var result = feature.FeatureDefinition{}
	result.Default.Value.Data = in.Value.Data // to be reviewed
	result.Default.Value.Type = in.Value.Type
	//TODO: complete other fields

	return result

}

func mapperFromFeatureDef(f *feature.FeatureDefinition) *grpcapipb.FeatureFullResponseFeatureDefinition {
	var result = new(grpcapipb.FeatureFullResponseFeatureDefinition)

	//TODO: implement mapper
	result.DefaultTtl = f.Default.TTL
	// .....other fields...

	return result
}

func (*ConfServer) Create(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureResponse, error) {
	var logic business.Business
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapperToFeatureDef(request.DefaultValue)

	result, err := logic.CreateFeature(vars, ruberConf)

	response := &grpcapipb.FeatureResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}

func (*ConfServer) Patch(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureResponse, error) {
	var logic business.Business
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapperToFeatureDef(request.DefaultValue)

	result, err := logic.PatchFeature(vars, ruberConf)

	response := &grpcapipb.FeatureResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}
func (*ConfServer) Delete(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, err := logic.DeleteFeature(vars)

	response := &grpcapipb.FeatureResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}

func (*FeatureServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureShortResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, value, typevalue, err := logic.GetFeatureOnlyValue(vars)

	valueStr := fmt.Sprintf("%v", value)

	response := &grpcapipb.FeatureShortResponse{
		Status: grpcapipb.StatusType(result),
		Value:  valueStr,
		Type:   typevalue,
	}
	return response, err
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

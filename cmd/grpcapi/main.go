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
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type ConfServer struct {
}
type FeatureServer struct {
}

func (*ConfServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureFullResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}
	result, content, err := logic.GetFeatureFull(vars)

	response := &grpcapipb.FeatureFullResponse{
		Status:  grpcapipb.StatusType(result),
		Feature: mapperFrom(content),
	}
	return response, err
}

//instead of mapping field by field, we've tagged fields with `json` tag. Every field has the same id in both
// structures. So we ca serialize and deserialize through json format.
func mapperTo(in *grpcapipb.FeatureDefinition) *feature.FeatureDefinition {

	var result = new(feature.FeatureDefinition)

	m := protojson.MarshalOptions{
		Indent:          "  ",
		EmitUnpopulated: true,
	}
	jsonBytes, err := m.Marshal(in)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(jsonBytes))

	err = json.Unmarshal(jsonBytes, result)
	if err != nil {
		panic(err)
	}
	return result

}

func mapperFrom(in *feature.FeatureDefinition) *grpcapipb.FeatureDefinition {
	var result = new(grpcapipb.FeatureDefinition)

	jsonBytes, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonBytes, result)
	if err != nil {
		panic(err)
	}
	return result
}

func (*ConfServer) Create(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureBasicResponse, error) {
	var logic business.Business
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapperTo(request.Feature)

	result, err := logic.CreateFeature(vars, *ruberConf)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}

func (*ConfServer) Patch(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureBasicResponse, error) {
	var logic business.Business
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapperTo(request.Feature)

	result, err := logic.PatchFeature(vars, *ruberConf)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}
func (*ConfServer) Delete(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureBasicResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, err := logic.DeleteFeature(vars)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}

func (*FeatureServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureShortResponse, error) {
	var logic business.Business
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, value, typevalue, err := logic.GetFeatureOnlyValue(vars)
	if err != nil {
		panic(err)
	}
	a, ok := value.(proto.Message)
	if !ok {
		panic("error casting")
	}
	b, err := anypb.New(a)
	if err != nil {
		panic(err)
	}
	response := &grpcapipb.FeatureShortResponse{
		Status: grpcapipb.StatusType(result),
		Value:  b,
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

package servers

import (
	"context"
	"encoding/json"

	"github.com/rubberyconf/rubberyconf/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/internal/feature"
	"github.com/rubberyconf/rubberyconf/internal/service"
	"google.golang.org/protobuf/encoding/protojson"
)

type ConfServer struct {
}

func (*ConfServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureFullResponse, error) {
	var logic service.Service
	name := request.FeatureName

	vars := map[string]string{"feature": name}
	result, content, err := logic.GetFeatureFull(ctx, vars)

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
	var logic service.Service
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapperTo(request.Feature)

	result, err := logic.CreateFeature(ctx, vars, *ruberConf)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}

func (*ConfServer) Patch(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureBasicResponse, error) {
	var logic service.Service
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapperTo(request.Feature)

	result, err := logic.PatchFeature(ctx, vars, *ruberConf)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}
func (*ConfServer) Delete(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureBasicResponse, error) {
	var logic service.Service
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, err := logic.DeleteFeature(ctx, vars)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}

package servers

import (
	"context"
	"encoding/json"

	"github.com/rubberyconf/rubberyconf/lib/application/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/lib/core/domain/feature"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/input"
	"google.golang.org/protobuf/encoding/protojson"
)

type ConfServer struct {
	service *input.IServiceFeature
}

func (me *ConfServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureFullResponse, error) {
	name := request.FeatureName

	vars := map[string]string{"feature": name}
	result, content, err := me.service.GetFeatureFull(ctx, vars)

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

func (me *ConfServer) Create(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureBasicResponse, error) {
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapperTo(request.Feature)

	result, err := me.service.CreateFeature(ctx, vars, *ruberConf)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}

func (me *ConfServer) Patch(ctx context.Context, request *grpcapipb.FeatureCreationRequest) (*grpcapipb.FeatureBasicResponse, error) {
	name := request.Name

	vars := map[string]string{"feature": name}

	ruberConf := mapperTo(request.Feature)

	result, err := me.service.PatchFeature(ctx, vars, *ruberConf)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}
func (me *ConfServer) Delete(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureBasicResponse, error) {
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, err := me.service.DeleteFeature(ctx, vars)

	response := &grpcapipb.FeatureBasicResponse{
		Status: grpcapipb.StatusType(result),
	}
	return response, err
}

func (me *ConfServer) SetService(srv *input.IServiceFeature) {
	me.service = srv
}

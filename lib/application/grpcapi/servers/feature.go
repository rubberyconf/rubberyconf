package servers

import (
	"context"

	"github.com/rubberyconf/rubberyconf/lib/application/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/lib/core/ports/input"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type FeatureServer struct {
	service *input.IServiceFeature
}

func (me *FeatureServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureShortResponse, error) {
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, value, typevalue, err := me.service.GetFeatureOnlyValue(ctx, vars)
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

func (me *FeatureServer) SetService(srv *input.IServiceFeature) {
	me.service = srv
}

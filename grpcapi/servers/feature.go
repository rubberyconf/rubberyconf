package servers

import (
	"context"

	"github.com/rubberyconf/rubberyconf/grpcapi/grpcapipb"
	"github.com/rubberyconf/rubberyconf/internal/service"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type FeatureServer struct {
}

func (*FeatureServer) Get(ctx context.Context, request *grpcapipb.FeatureIdRequest) (*grpcapipb.FeatureShortResponse, error) {
	var logic service.Service
	name := request.FeatureName

	vars := map[string]string{"feature": name}

	result, value, typevalue, err := logic.GetFeatureOnlyValue(ctx, vars)
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

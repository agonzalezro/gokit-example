package grpc

import (
	"context"

	"github.com/agonzalezro/alex-gokit-example/service/hiworld"
	"github.com/go-kit/kit/endpoint"
)

func MakeHiEndpoint(svc hiworld.Interface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		v := svc.Hi(req.Name)
		return Response{v}, nil
	}
}

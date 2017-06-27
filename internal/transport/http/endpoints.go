package http

import (
	"context"

	"github.com/agonzalezro/alex-gokit-example/service/hiworld"
	"github.com/go-kit/kit/endpoint"
)

func MakeHiEndpoint(svc hiworld.Interface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		if req.Name == "pepe" {
			return nil, ErrInvalidName // Just to have an easy example
		}
		v := svc.Hi(req.Name)
		return Response{v, ""}, nil
	}
}

func MakeByeEndpoint(svc hiworld.Interface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		return Response{svc.Bye(req.Name), ""}, nil
	}
}

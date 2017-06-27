package grpc

import (
	"context"

	"github.com/agonzalezro/alex-gokit-example/pb"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	V string `json:"v"`
}

// TODO: remote GRPC from the names here, it's part of the package already
func DecodeHiRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.HiRequest)
	return Request{Name: string(req.Name)}, nil
}

func EncodeHiResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(Response)
	return &pb.HiReply{V: string(resp.V)}, nil
}

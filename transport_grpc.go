package main

import (
	"context"

	"github.com/agonzalezro/hiworld/pb"
)

func DecodeGRPCHiRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.HiRequest)
	return hiWorldRequest{Name: string(req.Name)}, nil
}

func DecodeGRPCHiResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.HiReply)
	return hiWorldResponse{V: string(reply.V)}, nil
}

func EncodeGRPCHiResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(hiWorldResponse)
	return &pb.HiReply{V: string(resp.V)}, nil
}

func EncodeGRPCHiRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(hiWorldRequest)
	return &pb.HiRequest{Name: string(req.Name)}, nil
}

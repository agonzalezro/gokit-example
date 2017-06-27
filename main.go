package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/agonzalezro/hiworld/pb"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRequest  = errors.New("bad request")
	ErrBadRouting  = errors.New("inconsistent mapping between route and handler (programmer error)")
	ErrInvalidName = errors.New("the provided name is invalid")
)

func makeHiEndpoint(svc HiWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hiWorldRequest)
		if req.Name == "pepe" {
			return nil, ErrInvalidName
		}
		v := svc.Hi(req.Name)
		return hiWorldResponse{v, ""}, nil
	}
}

func makeByeEndpoint(svc HiWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hiWorldRequest)
		return hiWorldResponse{svc.Bye(req.Name), ""}, nil
	}
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc HiWorldService
	svc = hiWorldService{}
	svc = loggingMiddleware{logger, svc}

	hiEndpoint := makeHiEndpoint(svc)
	byeEndpoint := makeByeEndpoint(svc)

	go func() {
		options := []httptransport.ServerOption{
			httptransport.ServerErrorLogger(logger),
			httptransport.ServerErrorEncoder(encodeError),
		}

		hiHandler := httptransport.NewServer(
			hiEndpoint,
			decodeHiRequest,
			encodeResponse,
			options...,
		)

		byeHandler := httptransport.NewServer(
			byeEndpoint,
			decodeByeRequest,
			encodeResponse,
			options...,
		)

		r := mux.NewRouter()
		r.Path("/hi").Handler(hiHandler)
		r.Methods("GET").Path("/bye/{name}").Handler(byeHandler)

		var port = 8080
		hostAndPort := fmt.Sprintf(":%d", port)
		logger.Log("info", "HTTP listening on "+hostAndPort)
		logger.Log("err", http.ListenAndServe(hostAndPort, r))
	}()

	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Log("err", err)
		return
	}

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	srv := &grpcServer{
		hi: grpctransport.NewServer(
			hiEndpoint,
			DecodeGRPCHiRequest,
			EncodeGRPCHiResponse,
			options...,
		),
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, srv)

	logger.Log("info", "grpc listening on :8081")
	logger.Log("err", s.Serve(ln))
}

type grpcServer struct{ hi grpctransport.Handler }

func (s *grpcServer) Hi(ctx oldcontext.Context, req *pb.HiRequest) (*pb.HiReply, error) {
	_, rep, err := s.hi.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HiReply), nil
}

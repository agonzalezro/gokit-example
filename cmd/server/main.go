package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/agonzalezro/alex-gokit-example/internal/middleware"
	internalGRPCTransport "github.com/agonzalezro/alex-gokit-example/internal/transport/grpc"
	internalHTTPTransport "github.com/agonzalezro/alex-gokit-example/internal/transport/http"
	"github.com/agonzalezro/alex-gokit-example/pb"
	"github.com/agonzalezro/alex-gokit-example/service/hiworld"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc hiworld.Interface
	svc = hiworld.Service{}
	svc = middleware.Logging{logger, svc}

	go func() {
		options := []httptransport.ServerOption{
			httptransport.ServerErrorLogger(logger),
			httptransport.ServerErrorEncoder(internalHTTPTransport.EncodeError),
		}

		hiHandler := httptransport.NewServer(
			internalHTTPTransport.MakeHiEndpoint(svc),
			internalHTTPTransport.DecodeHiRequest,
			internalHTTPTransport.EncodeResponse,
			options...,
		)

		byeHandler := httptransport.NewServer(
			internalHTTPTransport.MakeByeEndpoint(svc),
			internalHTTPTransport.DecodeByeRequest,
			internalHTTPTransport.EncodeResponse,
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
			internalGRPCTransport.MakeHiEndpoint(svc),
			internalGRPCTransport.DecodeHiRequest,
			internalGRPCTransport.EncodeHiResponse,
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

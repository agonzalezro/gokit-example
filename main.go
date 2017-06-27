package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRequest  = errors.New("bad request")
	ErrBadRouting  = errors.New("inconsistent mapping between route and handler (programmer error)")
	ErrInvalidName = errors.New("the provided name is invalid")
)

func makeSalutateEndpoint(svc HiWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hiWorldRequest)
		if req.Name == "pepe" {
			return nil, ErrInvalidName
		}
		v := svc.Salutate(req.Name)
		return hiWorldResponse{v, ""}, nil
	}
}

func makeByeEndpoint(svc HiWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hiWorldRequest)
		return hiWorldResponse{svc.Bye(req.Name), ""}, nil
	}
}

func decodeSalutateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request hiWorldRequest // empty for now
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, ErrBadRequest
	}
	return request, nil
}

func decodeByeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		return nil, ErrBadRouting
	}
	return hiWorldRequest{Name: name}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc HiWorldService
	svc = hiWorldService{}
	svc = loggingMiddleware{logger, svc}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	salutateHandler := httptransport.NewServer(
		makeSalutateEndpoint(svc),
		decodeSalutateRequest,
		encodeResponse,
		options...,
	)

	byeHandler := httptransport.NewServer(
		makeByeEndpoint(svc),
		decodeByeRequest,
		encodeResponse,
		options...,
	)

	r := mux.NewRouter()
	r.Path("/hi").Handler(salutateHandler)
	r.Methods("GET").Path("/bye/{name}").Handler(byeHandler)

	// ws := new(restful.WebService)
	// ws.Route(ws.GET("/hi").To(salutateHandler))
	// ws.Route(ws.GET("/bye/{name}").To(byeHandler))
	// restful.Add(ws)

	var port = 8080
	hostAndPort := fmt.Sprintf(":%d", port)
	logger.Log("info", "Listening on "+hostAndPort)
	logger.Log("err", http.ListenAndServe(hostAndPort, r))
}

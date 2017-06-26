package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouting  = errors.New("inconsistent mapping between route and handler (programmer error)")
	ErrInvalidName = errors.New("The provided name is invalid")
)

type HiWorldService interface {
	Salutate(string) string
	Bye(string) string
}

type hiWorldService struct{}

func (hiWorldService) Salutate(name string) string {
	return fmt.Sprintf("Hi %s!", name)
}

func (hiWorldService) Bye(name string) string {
	return fmt.Sprintf("bye %s", name)
}

type hiWorldRequest struct {
	Name string `json:"name"`
}

type hiWorldResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

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
		return nil, err
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

type loggingMiddleware struct {
	logger log.Logger
	next   HiWorldService
}

func (mw loggingMiddleware) Salutate(name string) (output string) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "salutate",
			"input", name,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.next.Salutate(name)
	return output
}

func (mw loggingMiddleware) Bye(name string) (output string) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "bye",
			"input", name,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.next.Bye(name)
	return output
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

	logger.Log("err", http.ListenAndServe(":8080", r))
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrBadRouting:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

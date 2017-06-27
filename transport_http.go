package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func decodeHiRequest(_ context.Context, r *http.Request) (interface{}, error) {
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
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrBadRouting:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

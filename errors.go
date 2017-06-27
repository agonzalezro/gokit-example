package main

import (
	"context"
	"encoding/json"
	"net/http"
)

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

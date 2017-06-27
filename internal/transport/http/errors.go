package http

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest  = errors.New("bad request")
	ErrBadRouting  = errors.New("inconsistent mapping between route and handler (programmer error)")
	ErrInvalidName = errors.New("the provided name is invalid")
)

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

package main

import "fmt"

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

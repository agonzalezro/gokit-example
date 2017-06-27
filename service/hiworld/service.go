package hiworld

import "fmt"

type Interface interface {
	Hi(string) string
	Bye(string) string
}

type Service struct{}

func (Service) Hi(name string) string {
	return fmt.Sprintf("Hi %s!", name)
}

func (Service) Bye(name string) string {
	return fmt.Sprintf("bye %s", name)
}

package middleware

import (
	"time"

	"github.com/agonzalezro/alex-gokit-example/service/hiworld"
	"github.com/go-kit/kit/log"
)

type Logging struct {
	Logger log.Logger
	Next   hiworld.Interface
}

func (mw Logging) Hi(name string) (output string) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "hi",
			"input", name,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.Next.Hi(name)
	return output
}

func (mw Logging) Bye(name string) (output string) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "bye",
			"input", name,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.Next.Bye(name)
	return output
}

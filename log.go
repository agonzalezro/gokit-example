package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

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

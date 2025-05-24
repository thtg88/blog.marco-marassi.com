package middlewares

import (
	"net/http"
	"time"

	"github.com/thtg88/blog.marco-marassi.com/internal/dependencies"
	"github.com/thtg88/blog.marco-marassi.com/pkg/timer"
)

type TimerMiddeware struct {
	logger      dependencies.Logger
	nextHandler http.Handler
	time        timer.Timer
}

func NewTimerMiddleware(nextHandler http.Handler, logger dependencies.Logger, timer timer.Timer) *TimerMiddeware {
	return &TimerMiddeware{logger: logger, nextHandler: nextHandler, time: timer}
}

func (m *TimerMiddeware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		m.logger.Printf(
			"request to %s took: %v\n",
			r.URL.Path,
			m.time.Since(start),
		)
	}()

	m.nextHandler.ServeHTTP(w, r)
}

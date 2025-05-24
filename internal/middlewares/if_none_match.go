package middlewares

import (
	"net/http"

	"github.com/thtg88/blog.marco-marassi.com/internal/dependencies"
)

const IfNoneMatchHTTPHeaderName = "If-None-Match"

type IfNoneMatchMiddleware struct {
	etag        string
	logger      dependencies.Logger
	nextHandler http.Handler
}

func NewIfNoneMatchMiddleware(nextHandler http.Handler, etag string, logger dependencies.Logger) *IfNoneMatchMiddleware {
	return &IfNoneMatchMiddleware{
		etag:        etag,
		logger:      logger,
		nextHandler: nextHandler,
	}
}

func (m *IfNoneMatchMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(IfNoneMatchHTTPHeaderName) == m.etag {
		m.logger.Printf("request was cached with etag %s\n", m.etag)
		w.WriteHeader(http.StatusNotModified)
		return
	}

	m.nextHandler.ServeHTTP(w, r)
}

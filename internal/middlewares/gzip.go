package middlewares

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/thtg88/blog.marco-marassi.com/internal/dependencies"
	"github.com/thtg88/blog.marco-marassi.com/pkg/gzipwriter"
)

const (
	AcceptEncodingHTTPHeaderName   = "Accept-Encoding"
	ContentEncodingHTTPHeaderName  = "Content-Encoding"
	GzipEncoding                   = "gzip"
	GzipRequestNotSupportedLogLine = "request does not support gzip"
)

type GzipMiddleware struct {
	logger      dependencies.Logger
	nextHandler http.Handler
}

func NewGzipMiddleware(nextHandler http.Handler, logger dependencies.Logger) *GzipMiddleware {
	return &GzipMiddleware{
		logger:      logger,
		nextHandler: nextHandler,
	}
}

func (m *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get(AcceptEncodingHTTPHeaderName), GzipEncoding) {
		m.logger.Printf("request does not support gzip")
		m.nextHandler.ServeHTTP(w, r)
		return
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	gzipResponseWriter := gzipwriter.NewResponseWriter(gzipWriter, w)

	gzipResponseWriter.Header().Set(ContentEncodingHTTPHeaderName, GzipEncoding)
	m.nextHandler.ServeHTTP(gzipResponseWriter, r)
}

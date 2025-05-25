package middlewares

import (
	"fmt"
	"net/http"
)

const (
	CacheControlHTTPHeaderName        = "Cache-Control"
	CacheControlHTTPHeaderValueFormat = "max-age=%s, public"
)

var cachedPaths = map[string]string{
	"/style.css": "1209600",
}

type CacheControlMiddleware struct {
	nextHandler http.Handler
}

func NewCacheControlMiddleware(nextHandler http.Handler) *CacheControlMiddleware {
	return &CacheControlMiddleware{nextHandler: nextHandler}
}

func (m *CacheControlMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	durationSeconds, ok := cachedPaths[r.URL.Path]
	if ok {
		w.Header().Set(CacheControlHTTPHeaderName, fmt.Sprintf(CacheControlHTTPHeaderValueFormat, durationSeconds))
	}

	m.nextHandler.ServeHTTP(w, r)
}

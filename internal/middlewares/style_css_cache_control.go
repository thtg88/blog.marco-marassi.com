package middlewares

import "net/http"

const (
	StyleCSSPath                = "/style.css"
	CacheControlHTTPHeaderName  = "Cache-Control"
	CacheControlHTTPHeaderValue = "max-age=1209600, public"
)

type StyleCSSCacheControlMiddleware struct {
	nextHandler http.Handler
}

func NewStyleCSSCacheControlMiddleware(nextHandler http.Handler) *StyleCSSCacheControlMiddleware {
	return &StyleCSSCacheControlMiddleware{nextHandler: nextHandler}
}

func (m *StyleCSSCacheControlMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == StyleCSSPath {
		w.Header().Set(CacheControlHTTPHeaderName, CacheControlHTTPHeaderValue)
	}

	m.nextHandler.ServeHTTP(w, r)
}

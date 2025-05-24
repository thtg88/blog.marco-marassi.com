package handlers

import (
	"net/http"
)

// Static returns an http.Handler for static files.
// The handler supports ETags for improved caching.
func Static(fs http.Handler, etag string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ETag", etag)
		fs.ServeHTTP(w, r)
	}
}

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/thtg88/blog.marco-marassi.com/internal/dependencies"
	"github.com/thtg88/blog.marco-marassi.com/internal/handlers"
	"github.com/thtg88/blog.marco-marassi.com/internal/middlewares"
)

var deps *dependencies.Dependencies

func init() {
	deps = dependencies.Initialize()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Static(deps.FileServer, deps.ETag))

	// Middleware are processed in reverse order, the last one first.
	var handler http.Handler
	handler = middlewares.NewGzipMiddleware(mux, deps.Logger)
	handler = middlewares.NewCacheControlMiddleware(handler)
	handler = middlewares.NewIfNoneMatchMiddleware(handler, deps.ETag, deps.Logger)
	handler = middlewares.NewAllowedPathsMiddleware(handler, deps.Logger)
	handler = middlewares.NewTimerMiddleware(handler, deps.Logger, deps.Timer)

	deps.Logger.Printf("listening on port :%s", deps.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", deps.Port), handler)
	if errors.Is(err, http.ErrServerClosed) {
		deps.Logger.Println("server shutting down")
	} else if err != nil {
		deps.Logger.Fatal(err)
	}
}

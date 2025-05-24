package middlewares

import (
	"net/http"
	"slices"

	"github.com/thtg88/blog.marco-marassi.com/internal/dependencies"
)

type AllowedPathsMiddleware struct {
	logger      dependencies.Logger
	nextHandler http.Handler
}

var allowedPaths = []string{
	"/",
	"/favicons/apple-icon-57x57.png",
	"/favicons/apple-icon-60x60.png",
	"/favicons/apple-icon-72x72.png",
	"/favicons/apple-icon-76x76.png",
	"/favicons/apple-icon-114x114.png",
	"/favicons/apple-icon-120x120.png",
	"/favicons/apple-icon-144x144.png",
	"/favicons/apple-icon-152x152.png",
	"/favicons/apple-icon-180x180.png",
	"/favicons/android-icon-192x192.png",
	"/favicons/favicon-32x32.png",
	"/favicons/favicon-96x96.png",
	"/favicons/favicon-16x16.png",
	"/images/profile.png",
	"/posts/dns-not-working-after-deleting-a-linode",
	"/posts/dns-not-working-after-deleting-a-linode/",
	"/posts/laravel-passport-heroku-oauth-private-key-does-not-exist-or-is-not-readable",
	"/posts/laravel-passport-heroku-oauth-private-key-does-not-exist-or-is-not-readable/",
	"/posts/run-serverless-laravel-app-with-queue-workers-on-aws-lambda-using-bref",
	"/posts/run-serverless-laravel-app-with-queue-workers-on-aws-lambda-using-bref/",
	"/manifest.json",
	"/robots.txt",
	"/style.css",
}

func NewAllowedPathsMiddleware(nextHandler http.Handler, logger dependencies.Logger) *AllowedPathsMiddleware {
	return &AllowedPathsMiddleware{
		logger:      logger,
		nextHandler: nextHandler,
	}
}

func (m *AllowedPathsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !slices.Contains(allowedPaths, r.URL.Path) {
		m.logger.Printf("path not found %s\n", r.URL.Path)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	m.nextHandler.ServeHTTP(w, r)
}

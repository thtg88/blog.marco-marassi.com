package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/blog.marco-marassi.com/internal/middlewares"
)

func TestCacheControlMiddleware_ServeHTTP(t *testing.T) {
	type test struct {
		description          string
		path                 string
		expectedCacheControl string
	}

	tests := []test{
		{
			description:          "request to /style.css",
			path:                 "/style.css",
			expectedCacheControl: "max-age=1209600, public",
		},
		{
			description:          "any other request",
			path:                 "/",
			expectedCacheControl: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequest("GET", tc.path, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			middleware := middlewares.NewCacheControlMiddleware(handler)

			middleware.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, tc.expectedCacheControl, rr.Header().Get(middlewares.CacheControlHTTPHeaderName))
		})
	}
}

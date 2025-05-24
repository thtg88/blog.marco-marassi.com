package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/blog.marco-marassi.com/internal/middlewares"
	"github.com/thtg88/blog.marco-marassi.com/pkg/mocks"
)

func TestAllowedPathsMiddleware_ServeHTTP(t *testing.T) {
	type test struct {
		description      string
		path             string
		expectedCode     int
		expectedBody     string
		expectedLogLines []string
	}

	tests := []test{
		{
			description:  "not found path",
			path:         "/not-found",
			expectedCode: http.StatusNotFound,
			expectedBody: "not found\n",
			expectedLogLines: []string{
				"path not found /not-found\n",
			},
		},
		{
			description:  "/",
			path:         "/",
			expectedCode: http.StatusOK,
		},
		{
			description:  "/robots.txt",
			path:         "/robots.txt",
			expectedCode: http.StatusOK,
		},
		{
			description:  "/style.css",
			path:         "/style.css",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			logger := &mocks.LoggerMock{}

			req, err := http.NewRequest("GET", tc.path, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			middleware := middlewares.NewAllowedPathsMiddleware(nextHandler, logger)

			middleware.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, rr.Body.String())
			}

			if len(tc.expectedLogLines) == 0 {
				return
			}
			for _, expectedLogLine := range tc.expectedLogLines {
				assert.True(t, logger.HasLogged(expectedLogLine))
			}
		})
	}
}

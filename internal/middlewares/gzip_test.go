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

func TestGzipMiddleware_ServeHTTP(t *testing.T) {
	type test struct {
		description            string
		acceptEncodingHeader   string
		expectedEncodingHeader string
		expectedLogLine        string
	}

	tests := []test{
		{
			description:            "gzip encoding supported",
			acceptEncodingHeader:   middlewares.GzipEncoding,
			expectedEncodingHeader: middlewares.GzipEncoding,
		},
		{
			description:            "gzip encoding not supported",
			acceptEncodingHeader:   "",
			expectedEncodingHeader: "",
			expectedLogLine:        middlewares.GzipRequestNotSupportedLogLine,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequest("GET", "/", nil)
			require.NoError(t, err)

			req.Header.Add(middlewares.AcceptEncodingHTTPHeaderName, tc.acceptEncodingHeader)

			logger := &mocks.LoggerMock{}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			middleware := middlewares.NewGzipMiddleware(handler, logger)

			middleware.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, tc.expectedEncodingHeader, rr.Header().Get(middlewares.ContentEncodingHTTPHeaderName))

			if tc.expectedLogLine != "" {
				assert.True(t, logger.HasLogged(tc.expectedLogLine))
			}
		})
	}
}

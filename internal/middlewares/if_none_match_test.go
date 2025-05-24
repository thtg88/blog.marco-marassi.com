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

const etagTestValue = "etag"

func TestIfNoneMatchMiddleware_ServeHTTP(t *testing.T) {
	type test struct {
		description        string
		ifNoneMatchHeader  string
		expectedStatusCode int
		expectedLogLine    string
	}

	tests := []test{
		{
			description:        "cached request",
			ifNoneMatchHeader:  etagTestValue,
			expectedStatusCode: http.StatusNotModified,
			expectedLogLine:    "request was cached with etag etag\n",
		},
		{
			description:        "not cached request",
			ifNoneMatchHeader:  "not matching etag",
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			logger := &mocks.LoggerMock{}

			req, err := http.NewRequest("GET", "/", nil)
			require.NoError(t, err)

			req.Header.Add(middlewares.IfNoneMatchHTTPHeaderName, tc.ifNoneMatchHeader)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			middleware := middlewares.NewIfNoneMatchMiddleware(handler, etagTestValue, logger)

			middleware.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code)

			if tc.expectedLogLine != "" {
				assert.True(t, logger.HasLogged(tc.expectedLogLine))
			}
		})
	}
}

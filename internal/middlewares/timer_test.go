package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/blog.marco-marassi.com/internal/middlewares"
	"github.com/thtg88/blog.marco-marassi.com/pkg/mocks"
)

func TestTimerMiddleware_ServeHTTP(t *testing.T) {
	t.Run("timer logging", func(t *testing.T) {
		t.Parallel()

		logger := &mocks.LoggerMock{}
		time := mocks.NewTimerMock(time.Unix(0, 0))

		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		middleware := middlewares.NewTimerMiddleware(handler, logger, time)

		middleware.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.True(t, logger.HasLogged("request to / took: 1ms\n"))
	})
}

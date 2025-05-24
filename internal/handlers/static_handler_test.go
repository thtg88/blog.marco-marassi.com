package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/blog.marco-marassi.com/internal/handlers"
)

const etagTestValue = "etag"

func TestStatic(t *testing.T) {
	type test struct {
		description string
		path        string
	}

	tests := []test{
		{
			description: "successful request home",
			path:        "/",
		},
		{
			description: "successful request robots.txt",
			path:        "/robots.txt",
		},
		{
			description: "successful request style.css",
			path:        "/style.css",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			fs := http.FileServer(http.Dir("../../static"))

			req, err := http.NewRequest("GET", tc.path, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.Static(fs, etagTestValue))
			handler.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)
		})
	}
}

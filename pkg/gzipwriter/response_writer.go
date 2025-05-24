package gzipwriter

import (
	"io"
	"net/http"
)

type ResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func NewResponseWriter(writer io.Writer, responseWriter http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		Writer:         writer,
		ResponseWriter: responseWriter,
	}
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

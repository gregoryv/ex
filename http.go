// Package ex implements interfaces which writes output to stdout
package ex

import (
	"net/http"
	"fmt"
)

type BodyWriter struct {
	header http.Header
}

func NewBodyWriter() *BodyWriter {
	return &BodyWriter{
		header: http.Header{},
	}
}

// WriteHeader does nothing
func (w *BodyWriter) WriteHeader(v int)  {}

// Write the given byte slice to stdout
func (w *BodyWriter) Write(b []byte) (int, error)  {
	fmt.Printf("%s", string(b))
	return len(b), nil
}

func (w *BodyWriter) Header() http.Header  {
	return w.header
}

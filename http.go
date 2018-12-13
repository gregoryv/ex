// Package ex implements interfaces which writes output to stdout
package ex

import (
	"fmt"
	"net/http"
)

type BodyWriter struct {
	header http.Header
}

func NewBodyWriter() *BodyWriter {
	return &BodyWriter{header: http.Header{}}
}
func (w *BodyWriter) WriteHeader(v int)   {}
func (w *BodyWriter) Header() http.Header { return w.header }

// Write the given byte slice to stdout
func (w *BodyWriter) Write(b []byte) (int, error) {
	fmt.Printf("%s", string(b))
	return len(b), nil
}

type StatusWriter struct {
	header  http.Header
	written bool
}

func NewStatusWriter() *StatusWriter {
	return &StatusWriter{header: http.Header{}}
}

// Write the given value to stdout
func (w *StatusWriter) WriteHeader(v int) {
	fmt.Printf("%v", v)
	w.written = true
}
func (w *StatusWriter) Header() http.Header { return w.header }
func (w *StatusWriter) Write(b []byte) (int, error) {
	if !w.written {
		w.WriteHeader(http.StatusOK)
	}
	return len(b), nil
}

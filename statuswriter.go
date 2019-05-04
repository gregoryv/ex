// Package ex implements interfaces which writes output to stdout
package ex

import (
	"fmt"
	"net/http"
)

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

func StatusOf(r *http.Request, err error) (*StatusWriter, *http.Request) {
	if err != nil {
		panic(err)
	}
	return NewStatusWriter(), r
}

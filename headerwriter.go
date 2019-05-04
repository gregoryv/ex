// Package ex implements interfaces which writes output to stdout
package ex

import (
	"fmt"
	"net/http"
	"strings"
)

type HeaderWriter struct {
	header  http.Header
	written bool
}

func NewHeaderWriter() *HeaderWriter {
	return &HeaderWriter{header: http.Header{}}
}
func (w *HeaderWriter) WriteHeader(v int)   {}
func (w *HeaderWriter) Header() http.Header { return w.header }

// Write the given byte slice to stdout
func (w *HeaderWriter) Write(b []byte) (int, error) {
	if w.written {
		return len(b), nil
	}
	for k, v := range w.header {
		if len(v) == 1 {
			fmt.Printf("%s: %s\n", k, v[0])
			continue
		}
		fmt.Printf("%s: %s\n", k, strings.Join(v, "; "))
	}
	w.written = true
	return len(b), nil
}

func HeadersOf(r *http.Request, err error) (*HeaderWriter, *http.Request) {
	if err != nil {
		panic(err)
	}
	return NewHeaderWriter(), r
}

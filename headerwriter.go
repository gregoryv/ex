// Package ex implements interfaces which writes output to stdout
package ex

import (
	"fmt"
	"net/http"
	"sort"
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
	sorted := make([]string, len(w.header))
	var i int
	for k, _ := range w.header {
		sorted[i] = k
		i++
	}
	sort.Sort(sort.StringSlice(sorted))
	for _, k := range sorted {
		v := w.header[k]
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

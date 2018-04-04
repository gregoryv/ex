// Package exio mocks varios write interfaces for eazy test examples.
package exio

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

func (w *BodyWriter) WriteHeader(v int)  {}

func (w *BodyWriter) Write(b []byte) (int, error)  {
	fmt.Printf("%s", string(b))
	return len(b), nil
}

func (w *BodyWriter) Header() http.Header  {
	return w.header
}

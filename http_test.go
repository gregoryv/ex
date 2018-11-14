package ex

import (
	"testing"
)

var bodyWriter = NewBodyWriter()
var statusWriter = NewStatusWriter()

func ExampleBodyWriter_Write() {
	bodyWriter.Write([]byte("Hello, world!"))
	// output:
	// Hello, world!
}

func ExampleBodyWriter_WriteHeader() {
	bodyWriter.WriteHeader(200)
	// output:
}

func TestBodyWriter_Header(t *testing.T) {
	if bodyWriter.Header() == nil {
		t.Fail()
	}
}

func ExampleStatusWriter_WriteHeader() {
	statusWriter.WriteHeader(200)
	// output:
	// 200
}
func ExampleStatusWriter_Write() {
	w := NewStatusWriter()
	w.Write([]byte("Hello, world!"))
	// output:
	// 200
}

func TestStatusWriter_Header(t *testing.T) {
	if statusWriter.Header() == nil {
		t.Fail()
	}
}

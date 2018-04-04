package ex

import (
	"testing"
)

func ExampleBodyWriter_Write() {
	w := NewBodyWriter()
	w.Write([]byte("Hello, world!"))
	// output:
	// Hello, world!
}

func ExampleBodyWriter_WriteHeader() {
	w := NewBodyWriter()
	w.WriteHeader(200)
	// output:
}

func TestBodyWriter_Header(t *testing.T) {
	w := NewBodyWriter()
	if w.Header() == nil {
		t.Fail()
	}
}

package ex

import (
	"fmt"
	"net/http"
	"testing"
)

var bodyWriter = NewBodyWriter()

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

func ExampleBodyOf() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", "world")
	}
	handler(BodyOf(http.NewRequest("GET", "/", nil)))
	//output:
	//Hello, world!
}

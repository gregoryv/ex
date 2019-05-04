package ex

import (
	"fmt"
	"net/http"
	"testing"
)

var headerWriter = NewHeaderWriter()

func ExampleHeaderWriter_Write() {
	headerWriter.Write([]byte("body not visible"))
	// output:
}

func ExampleHeaderWriter_WriteHeader() {
	w := NewHeaderWriter()
	w.Header().Add("Content-type", "text/plain")
	w.Header().Add("Accept", "*")
	w.Header().Add("Accept", "application/json")
	w.Write([]byte("body not visible"))
	// output:
	// Content-Type: text/plain
	// Accept: *; application/json
}

func TestHeaderWriter_Header(t *testing.T) {
	if headerWriter.Header() == nil {
		t.Fail()
	}
}

func ExampleHeadersOf() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", "world")
	}
	handler(HeadersOf(http.NewRequest("GET", "/", nil)))
	//output:
}

func TestHeadersof_panic(t *testing.T) {
	defer func() {
		e := recover()
		if e == nil {
			t.Error("Expected a panic")
		}
	}()
	HeadersOf(nil, fmt.Errorf("bad stuff"))
}

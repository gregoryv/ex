package ex

import (
	"net/http"
	"testing"
)

var statusWriter = NewStatusWriter()

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

func ExampleStatusOf() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	handler(StatusOf(http.NewRequest("GET", "/", nil)))
	//output:
	//200
}

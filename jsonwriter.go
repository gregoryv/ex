package ex

import (
	"fmt"
	"net/http"
	"strings"
)

// Very basic Json tidier
type JsonWriter struct {
	header http.Header
}

func NewJsonWriter() *JsonWriter          { return &JsonWriter{header: http.Header{}} }
func (w *JsonWriter) WriteHeader(v int)   {}
func (w *JsonWriter) Header() http.Header { return w.header }

// Write the given byte slice to stdout
func (w *JsonWriter) Write(b []byte) (int, error) {
	indent := 0
	inside := false
	for _, b := range b {
		switch b {
		case '"':
			inside = !inside
			fmt.Print(string(b))
		case ',':
			if inside {
				fmt.Print(string(b))
				continue
			}
			fmt.Print(",\n", strings.Repeat(" ", indent))
		case '{':
			indent += 4
			fmt.Print("{\n", strings.Repeat(" ", indent))
		case '}':
			indent -= 4
			fmt.Print("\n", strings.Repeat(" ", indent), "}")
		default:
			fmt.Print(string(b))
		}
	}
	return len(b), nil
}

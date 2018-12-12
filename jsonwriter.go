package ex

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Very basic Json tidier
type JsonWriter struct {
	header http.Header
	indent int
	inside bool
}

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{
		header: http.Header{},
	}
}
func (w *JsonWriter) WriteHeader(v int)   {}
func (w *JsonWriter) Header() http.Header { return w.header }

// Write the given byte slice to stdout
func (w *JsonWriter) Write(b []byte) (int, error) {
	return jwrite(w, os.Stdout, b)
}

func jwrite(jw *JsonWriter, out io.Writer, b []byte) (int, error) {
	for _, b := range b {
		switch b {
		case '"':
			jw.inside = !jw.inside
			fmt.Fprint(out, string(b))
		case ',':
			if jw.inside {
				fmt.Fprint(out, string(b))
				continue
			}
			fmt.Fprintf(out, ",\n%s", jw.tab())
		case '{':
			jw.indent += 4
			fmt.Fprintf(out, "{\n%s", jw.tab())
		case '}':
			jw.indent -= 4
			fmt.Fprintf(out, "\n%s}", jw.tab())
		default:
			fmt.Fprint(out, string(b))
		}
	}
	return len(b), nil
}

func (jw *JsonWriter) tab() string {
	return strings.Repeat(" ", jw.indent)
}

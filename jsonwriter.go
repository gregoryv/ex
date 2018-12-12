package ex

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// JsonWriter formats incomming json and implements http.ResponseWriter.
type JsonWriter struct {
	header http.Header
	indent int  // current indent size in spaces
	inside bool // wether inside value or not
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

// jwrite writes the given byte slice to out while tidying the json.
func jwrite(state *JsonWriter, out io.Writer, b []byte) (int, error) {
	for _, b := range b {
		switch b {
		case '"':
			state.inside = !state.inside
			fmt.Fprint(out, string(b))
		case ',':
			if state.inside {
				fmt.Fprint(out, string(b))
				continue
			}
			fmt.Fprintf(out, ",\n%s", state.tab())
		case '{':
			state.indent += 4
			fmt.Fprintf(out, "{\n%s", state.tab())
		case '}':
			state.indent -= 4
			fmt.Fprintf(out, "\n%s}", state.tab())
		default:
			fmt.Fprint(out, string(b))
		}
	}
	return len(b), nil
}

func (w *JsonWriter) tab() string {
	return strings.Repeat(" ", w.indent)
}

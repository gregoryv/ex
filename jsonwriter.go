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
	header      http.Header
	indent      int  // current indent size in spaces
	escaped     bool // wether last character was \
	byteHandler func(byte) string
}

func NewJsonWriter() *JsonWriter {
	w := &JsonWriter{
		header: http.Header{},
	}
	w.byteHandler = w.outside
	return w
}
func (w *JsonWriter) WriteHeader(v int)   {}
func (w *JsonWriter) Header() http.Header { return w.header }

// Write the given byte slice to stdout
func (w *JsonWriter) Write(b []byte) (int, error) {
	return w.fwrite(os.Stdout, b)
}

// jwrite writes the given byte slice to out while tidying the json.
func (w *JsonWriter) fwrite(out io.Writer, b []byte) (int, error) {
	for _, b := range b {
		fmt.Fprint(out, w.byteHandler(b))
	}
	return len(b), nil
}

// inside converts bytes within names or values in json
func (w *JsonWriter) inside(b byte) string {
	switch {
	case b == '\\':
		w.escaped = true
	case b == '"' && !w.escaped:
		w.byteHandler = w.outside
	default:
		w.escaped = false
	}
	return string(b)
}

// outside converts bytes that are outside of name or values in json
func (w *JsonWriter) outside(b byte) string {
	switch b {
	case '"':
		w.byteHandler = w.inside
	case ',':
		return ",\n" + w.tab()
	case '{':
		w.indent += 4
		return "{\n" + w.tab()
	case '}':
		w.indent -= 4
		return "\n" + w.tab() + "}"
	}
	return string(b)
}

func (w *JsonWriter) tab() string {
	return strings.Repeat(" ", w.indent)
}

package ex

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

func Fprint(w io.Writer, any ...interface{}) {
	for _, e := range any {
		switch e := e.(type) {
		case http.Header:
			fprint(os.Stdout, e)
		}
	}
}

func Print(any ...interface{}) {
	Fprint(os.Stdout, any...)
}

func fprint(w io.Writer, header http.Header) {
	keys := make([]string, 0)
	for k := range header {
		keys = append(keys, k)
	}
	sort.StringSlice(keys).Sort()
	for _, k := range keys {
		v := header[k]
		for _, v := range v {
			fmt.Fprintf(w, "%v: %v\n", k, v)
		}
	}
}

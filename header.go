package ex

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

// Print writes any element to stdout.
func Print(any ...interface{}) {
	Tfprint(asIs, os.Stdout, any...)
}

func Tprint(tidy TidyFunc, any ...interface{}) {
	Tfprint(tidy, os.Stdout, any...)
}

func Tfprint(tidy TidyFunc, w io.Writer, any ...interface{}) {
	for _, e := range any {
		switch e := e.(type) {
		case http.Header:
			printHeaders(tidy, os.Stdout, e)
		}
	}
}

func printHeaders(tidy TidyFunc, w io.Writer, header http.Header) {
	keys := make([]string, 0)
	for k := range header {
		keys = append(keys, k)
	}
	sort.StringSlice(keys).Sort()
	for _, k := range keys {
		v := header[k]
		for _, v := range v {
			fmt.Fprintf(w, "%v: %v\n", k, tidy(v))
		}
	}
}

type TidyFunc func(string) string

func asIs(v string) string { return v }

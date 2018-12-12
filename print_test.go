package ex

import (
	"net/http"
	"strings"
)

func ExamplePrint() {
	h := http.Header{}
	h.Add("Content-Type", "text/plain")
	h.Add("Abbe", "one")
	h.Add("Abbe", "two")
	Print(h)
	// output:
	// Abbe: one
	// Abbe: two
	// Content-Type: text/plain
}

func ExampleTprint() {
	h := http.Header{}
	h.Add("Content-Type", "text/plain")
	h.Add("Abbe", "one")
	h.Add("Abbe", "two")
	tidy := func(v string) string {
		return strings.Replace(v, "one", "zero", -1)
	}
	Tprint(tidy, h)
	// output:
	// Abbe: zero
	// Abbe: two
	// Content-Type: text/plain
}

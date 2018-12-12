package ex

import (
	"net/http"
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

package ex

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

// DumpParts prints each part of a multipart message
func DumpParts(msg io.Reader, boundary string) {
	mr := multipart.NewReader(msg, boundary)
	jsonWriter := NewJsonWriter()
	for p, _ := mr.NextPart(); p != nil; p, _ = mr.NextPart() {
		if isJson(p.Header["Content-Type"]) {
			io.Copy(jsonWriter, p)
			fmt.Fprint(jsonWriter, "\n")
		} else {
			io.Copy(os.Stdout, p)
		}
	}
}

func isJson(contentType []string) bool {
	for _, t := range contentType {
		if strings.Contains(t, "application/json") {
			return true
		}
	}
	return false
}

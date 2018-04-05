package ex

import (
	"io"
	"mime/multipart"
	"os"
)

// DumpParts prints each part of a multipart message
func DumpParts(msg io.Reader, boundary string) {
	mr := multipart.NewReader(msg, boundary)
	for p, _ := mr.NextPart(); p != nil; p, _ = mr.NextPart() {
		io.Copy(os.Stdout, p)
	}
}

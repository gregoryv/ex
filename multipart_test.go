package ex

import (
	"bytes"
	"mime/multipart"
	"net/textproto"
)

func ExampleDumpParts() {
	buf := bytes.NewBufferString("")
	body := multipart.NewWriter(buf)
	body.SetBoundary("--X")

	w, _ := body.CreatePart(textproto.MIMEHeader{})
	w.Write([]byte("Hello"))

	w, _ = body.CreatePart(textproto.MIMEHeader{})
	w.Write([]byte(", world!"))

	DumpParts(buf, "--X")
	// output:
	// Hello, world!
}

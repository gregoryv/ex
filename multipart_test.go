package ex

import (
	"bytes"
	"mime/multipart"
	"net/textproto"
)

func ExampleDumpParts_with_nice_json() {
	buf := bytes.NewBufferString("")
	body := multipart.NewWriter(buf)
	body.SetBoundary("--X")

	header := textproto.MIMEHeader{}
	header["Content-Type"] = []string{"application/json"}
	w, _ := body.CreatePart(header)
	w.Write([]byte(`{"text":"hello","more":"world"}`))

	w, _ = body.CreatePart(textproto.MIMEHeader{})
	w.Write([]byte("next part"))

	DumpParts(buf, "--X")
	// output:
	// {
	//     "text":"hello",
	//     "more":"world"
	// }
	// next part
}

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

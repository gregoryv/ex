package ex

import (
	"bytes"
	"io"
	"testing"
)

var jsonWriter = NewJsonWriter()

func ExampleJsonWriter_Write() {
	json := `{"name":"John","car":{"model":"x,2","plate":"abc\"123"}}`
	reader := bytes.NewBufferString(json)
	io.Copy(jsonWriter, reader)
	// output:
	// {
	//     "name":"John",
	//     "car":{
	//         "model":"x,2",
	//         "plate":"abc\"123"
	//     }
	// }
}

func ExampleJsonWriter_WriteHeader() {
	jsonWriter.WriteHeader(200)
	// output:
}

func TestJsonWriter_Header(t *testing.T) {
	if jsonWriter.Header() == nil {
		t.Fail()
	}
}

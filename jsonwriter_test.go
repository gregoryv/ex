package ex

import (
	"testing"
)

var jsonWriter = NewJsonWriter()

func ExampleJsonWriter_Write() {
	jsonWriter.Write([]byte(`{"name":"John","car":{"model":"x,2","plate":"abc123"}}`))
	// output:
	// {
	//     "name":"John",
	//     "car":{
	//         "model":"x,2",
	//         "plate":"abc123"
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

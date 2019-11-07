package ex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func ExampleJsonWriter_Write_specific_output() {
	json := `{"name":"John","car":{"model":"x,2","plate":"abc\"123"}}`
	reader := bytes.NewBufferString(json)
	jsonWriter := NewJsonWriter()
	jsonWriter.Out = os.Stdout
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

func ExampleJsonOf() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(
			struct {
				Message string
			}{"Hello, World!"})
	}
	handler(JsonOf(http.NewRequest("GET", "/", nil)))
	// output:
	// {
	//     "Message":"Hello, World!"
	// }
}

func ExampleJsonWriter_WriteTo() {
	json := `{"name":"John","car":{"model":"x,2","plate":"abc\"123"}}`
	buf := bytes.NewBufferString("")
	jsonWriter.WriteTo(buf, []byte(json))
	fmt.Print(buf.String())
	// output:
	// {
	//     "name":"John",
	//     "car":{
	//         "model":"x,2",
	//         "plate":"abc\"123"
	//     }
	// }
}

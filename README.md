ARCHIVED! Moved to https://sogvin.com/ex

[ex](https://godoc.org/github.com/gregoryv/ex) - implements interfaces which write output to stdout

This is useful for writing short Examples in go, eg. when testing output
from http handlers.

## Examples

Writers implement the http.ResponseWriter interface

### JsonWriter

    r := bytes.NewBufferString(`{"name":"John","car":{"model":"x,2","plate":"abc123"}}`)
	io.Copy(jsonWriter, r)
	
Output:

    {
        "name":"John",
        "car":{
            "model":"x,2",
            "plate":"abc123"
        }
    }


### StatusWriter

    w := NewStatusWriter()
    w.Write([]byte("Hello, world!"))
	
Output:

    200
	
	
### BodyWriter

    bodyWriter.Write([]byte("Hello, world!"))
	
Output:

    Hello, world!
	
	


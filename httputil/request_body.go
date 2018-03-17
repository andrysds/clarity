package httputil

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Read and unmarshal incoming http request body.
func RequestBody(bodyReader io.Reader) (interface{}, error) {
	var data interface{}
	body, _ := ioutil.ReadAll(bodyReader)
	err := json.Unmarshal(body, &data)
	return data, err
}

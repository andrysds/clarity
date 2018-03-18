package httputil

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// RequestBody reads and unmarshals incoming http request body.
func RequestBody(bodyReader io.Reader) (interface{}, error) {
	var data interface{}
	body, _ := ioutil.ReadAll(bodyReader)
	err := json.Unmarshal(body, &data)
	return data, err
}

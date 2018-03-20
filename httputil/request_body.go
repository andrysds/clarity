package httputil

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// RequestBody reads and unmarshals incoming http request body.
func RequestBody(bodyReader io.Reader, data interface{}) error {
	body, _ := ioutil.ReadAll(bodyReader)
	err := json.Unmarshal(body, data)
	return err
}

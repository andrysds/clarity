package clarity

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func RequestBody(bodyReader io.Reader) (interface{}, error) {
	var data interface{}
	body, _ := ioutil.ReadAll(bodyReader)
	err := json.Unmarshal(body, &data)
	return data, err
}

package clarity

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestRequestBody(t *testing.T) {
	var request *http.Request
	var body []byte
	var result interface{}
	var err error
	data := map[string]interface{}(
		map[string]interface{}{"key": "value"},
	)

	body, _ = json.Marshal(data)
	request, _ = http.NewRequest("GET", "http://localhost", bytes.NewBuffer(body))
	result, err = RequestBody(request.Body)
	if !reflect.DeepEqual(result, data) {
		t.Error("`result` should be equals `data`")
	}
	if err != nil {
		t.Error("`err` should be nil")
	}

	body = []byte("errror")
	request, _ = http.NewRequest("GET", "http://localhost", bytes.NewBuffer(body))
	result, err = RequestBody(request.Body)
	if reflect.DeepEqual(result, data) {
		t.Error("`result` should not be equals `data`")
	}
	if err == nil {
		t.Error("`err` should not be nil")
	}
}

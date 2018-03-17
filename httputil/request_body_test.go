package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestBody(t *testing.T) {
	var request *http.Request
	var body []byte
	var result interface{}
	var err error
	data := map[string]interface{}(
		map[string]interface{}{
			"key": "value",
		},
	)

	// error on reading body
	body = []byte("errror")
	request, _ = http.NewRequest("GET", "http://localhost", bytes.NewBuffer(body))
	result, err = RequestBody(request.Body)
	assert.NotEqual(t, result, data)
	assert.NotNil(t, err)

	// normal case
	body, _ = json.Marshal(data)
	request, _ = http.NewRequest("GET", "http://localhost", bytes.NewBuffer(body))
	result, err = RequestBody(request.Body)
	assert.Equal(t, result, data)
	assert.Nil(t, err)
}

func ExampleRequestBody() {
	data := map[string]interface{}(
		map[string]interface{}{
			"key": "value",
		},
	)
	bodyJSON, _ := json.Marshal(data)
	request, _ := http.NewRequest("GET", "http://localhost", bytes.NewBuffer(bodyJSON))
	body, err := RequestBody(request.Body)
	fmt.Println(body, err)
}

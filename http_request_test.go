package clarity

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestParamTest(t *testing.T) {
	var result interface{}
	URL, _ := url.Parse("http://localhost?key=value&number=1&long_number=10&array[]=elem1,elem2")
	wrongURL, _ := url.Parse("http://localhost?key=value&number=one&long_number=ten&array[]=elem1,elem2")
	defaults := map[string]interface{}{
		"s":   "default",
		"i":   0,
		"i64": int64(0),
		"a":   []string{"default"},
	}

	// context: missing param
	// string param
	result = RequestParam(URL, "miss", defaults["s"]).(string)
	assert.Equal(t, result, defaults["s"])

	// int param
	result = RequestParam(URL, "miss", defaults["i"]).(int)
	assert.Equal(t, result, defaults["i"])

	// int64 param
	result = RequestParam(URL, "miss", defaults["i64"]).(int64)
	assert.Equal(t, result, defaults["i64"])

	// array param
	result = RequestParam(URL, "miss", defaults["a"]).([]string)
	assert.Equal(t, result, defaults["a"])

	// context: existing param
	// string param
	result = RequestParam(URL, "key", defaults["s"]).(string)
	assert.Equal(t, result, "value")

	// int param
	result = RequestParam(URL, "number", defaults["i"]).(int)
	assert.Equal(t, result, 1)

	// wrong int param
	result = RequestParam(wrongURL, "number", defaults["i"]).(int)
	assert.Equal(t, result, 0)

	// int64 param
	result = RequestParam(URL, "long_number", defaults["i64"]).(int64)
	assert.Equal(t, result, int64(10))

	// wrong int64 param
	result = RequestParam(wrongURL, "long_number", defaults["i64"]).(int64)
	assert.Equal(t, result, int64(0))

	// array param
	result = RequestParam(URL, "array", defaults["a"]).([]string)
	assert.Equal(t, result, []string{"elem1", "elem2"})
}

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
	err = RequestBody(request.Body, &result)
	assert.NotEqual(t, result, data)
	assert.NotNil(t, err)

	// normal case
	body, _ = json.Marshal(data)
	request, _ = http.NewRequest("GET", "http://localhost", bytes.NewBuffer(body))
	err = RequestBody(request.Body, &result)
	assert.Equal(t, result, data)
	assert.Nil(t, err)
}

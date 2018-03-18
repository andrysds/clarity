package httputil

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestParamTest(t *testing.T) {
	var result interface{}
	URL, _ := url.Parse("http://localhost?key=value&number=1&array[]=elem1,elem2")
	defaults := map[string]interface{}{
		"s": "default",
		"i": 0,
		"a": []string{"default"},
	}

	// context: missing param
	// string param
	result = RequestParam(URL, "miss", defaults["s"]).(string)
	assert.Equal(t, result, defaults["s"])

	// int param
	result = RequestParam(URL, "miss", defaults["i"]).(int)
	assert.Equal(t, result, defaults["i"])

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

	// array param
	result = RequestParam(URL, "array", defaults["a"]).([]string)
	assert.Equal(t, result, []string{"elem1", "elem2"})
}

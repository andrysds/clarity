package clarity

import (
	"net/url"
	"reflect"
	"testing"
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
	if result != defaults["s"] {
		t.Error("result:", result)
		t.Error("result should be equals ", defaults["s"])
	}

	// int param
	result = RequestParam(URL, "miss", defaults["i"]).(int)
	if result != defaults["i"] {
		t.Error("result:", result)
		t.Error("result should be equals ", defaults["i"])
	}

	// array param
	result = RequestParam(URL, "miss", defaults["a"]).([]string)
	if !reflect.DeepEqual(result, defaults["a"]) {
		t.Error("result:", result)
		t.Error("result should be equals ", defaults["a"])
	}

	// context: existing param
	// string param
	result = RequestParam(URL, "key", defaults["s"]).(string)
	if result != "value" {
		t.Error("result:", result)
		t.Error("result should be equals \"value\"")
	}

	// int param
	result = RequestParam(URL, "number", defaults["i"]).(int)
	if result != 1 {
		t.Error("result:", result)
		t.Error("result should be equals 1")
	}

	// array param
	result = RequestParam(URL, "array", defaults["a"]).([]string)
	if !reflect.DeepEqual(result, []string{"elem1", "elem2"}) {
		t.Error("result:", result)
		t.Error("result should be equals []string{'elem1', 'elem2'}")
	}
}

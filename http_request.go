package clarity

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// RequestParam gets parameter from incoming http request url and refine it with a default value.
func RequestParam(URL *url.URL, key string, defaultValue interface{}) interface{} {
	switch defaultValue.(type) {
	case []string:
		key = key + "[]"
	}

	value := URL.Query().Get(key)
	if value != "" {
		switch defaultValue.(type) {
		case string:
			return value
		case int:
			intValue, _ := strconv.Atoi(value)
			return intValue
		case int64:
			int64Value, _ := strconv.ParseInt(value, 10, 64)
			return int64Value
		case []string:
			return strings.Split(value, ",")
		}
	}
	return defaultValue
}

// RequestBody reads and unmarshals incoming http request body.
func RequestBody(bodyReader io.Reader, data interface{}) error {
	body, _ := ioutil.ReadAll(bodyReader)
	err := json.Unmarshal(body, data)
	return err
}

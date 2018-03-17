package clarity

import (
	"net/url"
	"strconv"
	"strings"
)

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
		case []string:
			return strings.Split(value, ",")
		}
	}
	return defaultValue
}

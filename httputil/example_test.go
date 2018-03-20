package httputil_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/andrysds/clarity/httputil"
)

func ExampleBasicAuthorizer() {
	authorizer := httputil.NewBasicAuthorizer(
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
	)
	var header http.Header
	header.Set("Authorization", "Basic Y2xhcml0eTpjbGFyaXR5")
	err := authorizer.Authorize(header)
	fmt.Println(err)
}

func ExampleRequestBody() {
	data := map[string]interface{}(
		map[string]interface{}{
			"key": "value",
		},
	)
	bodyJSON, _ := json.Marshal(data)
	request, _ := http.NewRequest("POST", "http://localhost", bytes.NewBuffer(bodyJSON))
	var body interface{}
	err := httputil.RequestBody(request.Body, &body)
	fmt.Println(body, err)
}

func ExampleRequestParam() {
	URL, _ := url.Parse("http://localhost?key=value")
	key := httputil.RequestParam(URL, "key", "default").(string)
	fmt.Println(key)
}

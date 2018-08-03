package clarity_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/andrysds/clarity"
	_ "github.com/go-sql-driver/mysql"
)

func ExampleBasicAuthorizer() {
	authorizer := clarity.NewBasicAuthorizer(
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
	)
	var header http.Header
	header.Set("Authorization", "Basic Y2xhcml0eTpjbGFyaXR5")
	err := authorizer.Authorize(header)
	fmt.Println(err)
}

func ExamplePanicIfError() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	clarity.PanicIfError(err, "error on connecting mysql")
	return db
}

func ExamplePrintIfError() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	err := json.Unmarshal(byt, &dat)
	clarity.PrintIfError(err, "error on json unmarshal")
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
	err := clarity.RequestBody(request.Body, &body)
	fmt.Println(body, err)
}

func ExampleRequestParam() {
	URL, _ := url.Parse("http://localhost?key=value")
	key := clarity.RequestParam(URL, "key", "default").(string)
	fmt.Println(key)
}

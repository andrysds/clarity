# Clarity
[![GoDoc](https://godoc.org/github.com/andrysds/clarity?status.svg)](https://godoc.org/github.com/andrysds/clarity) [![Build Status](https://travis-ci.org/andrysds/clarity.svg)](https://travis-ci.org/andrysds/clarity) [![codecov](https://codecov.io/gh/andrysds/clarity/branch/master/graph/badge.svg)](https://codecov.io/gh/andrysds/clarity) [![Go Report Card](https://goreportcard.com/badge/github.com/andrysds/clarity)](https://goreportcard.com/report/github.com/andrysds/clarity)

Random but useful and essential utilites for your Go/Golang project.

## Installing

```
go get -u github.com/andrysds/clarity
```

## Packages

### errutil
[![GoDoc](https://godoc.org/github.com/andrysds/clarity/errutil?status.svg)](https://godoc.org/github.com/andrysds/clarity/errutil)

Package `errutil` provides a set of utilities for handling error in Go. For example, calls Panic if there is error on initiating you project's database.

Example:
```go
import (
	"database/sql"

	"github.com/andrysds/clarity/errutil"
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	// Calls Panic if there is error on connecting mysql.
	errutil.PanicIfError(err, "error on connecting mysql")
	return db
}
```

### httputil
[![GoDoc](https://godoc.org/github.com/andrysds/clarity/httputil?status.svg)](https://godoc.org/github.com/andrysds/clarity/httputil)

Package `httputil` provides a set of utilities for handling http (request, authorization) in Go.

Example:
```go
import (
  "net/http"
  "os"
  "fmt"

  "github.com/andrysds/clarity/httputil"
)

func HandleRequest(w *http.ResponseWriter, r *http.Request) {
  authorizer := httputil.NewBasicAuthorizer(
      os.Getenv("USERNAME"),
      os.Getenv("PASSWORD"),
  )
  err := authorizer.Authorize(r.Header)
  if err != nil {
    fmt.Println(err)
    return
  }

  paramKey := httputil.RequestParam(r.URL, "key", "default").(string)

  requestBody, err := httputil.RequestBody(r.Body)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Param 'key': ", paramKey)
  fmt.Println("Request Body: ", requestBody)
}
```

## Contributing

Feel free to submit pull requests.

## Versioning

For the versions available, see the [tags on this repository](https://github.com/andrysds/clarity/tags).

## Author

**Andrys Daniel Silalahi** - [andrysds](https://github.com/andrysds)

## License

This project is licensed under the MIT License ([LICENSE.md](https://github.com/andrysds/clarity/blob/master/LICENSE))

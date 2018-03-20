// Package httputil provides a set of utilities for handling http (request, authorization) in Go.
//
// Example Usage
//
// The following is a complete example using httputil:
//   import (
//     "net/http"
//     "os"
//     "fmt"
//
//     "github.com/andrysds/clarity/httputil"
//   )
//
//   func HandleRequest(w *http.ResponseWriter, r *http.Request) {
//     authorizer := httputil.NewBasicAuthorizer(
//         os.Getenv("USERNAME"),
//         os.Getenv("PASSWORD"),
//     )
//     err := authorizer.Authorize(r.Header)
//     if err != nil {
//       fmt.Println(err)
//       return
//     }
//
//     paramKey := httputil.RequestParam(r.URL, "key", "default").(string)
//
//     var requestBody interface{}
//     err = httputil.RequestBody(r.Body, &requestBody)
//     if err != nil {
//       fmt.Println(err)
//       return
//     }
//
//     fmt.Println("Param 'key': ", paramKey)
//     fmt.Println("Request Body: ", requestBody)
//   }
package httputil

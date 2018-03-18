// Package errutil provides a set of utilities for handling error in Go.
// For example, calls Panic if there is error on initiating you project's database.
package errutil

import "fmt"

// PanicIfError calls panic if the given error is not nil
func PanicIfError(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", message, err))
	}
}

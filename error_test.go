package clarity

import (
	"errors"
	"testing"
)

func TestPanicIfError(t *testing.T) {
	var result bool
	message := "error message"

	// err is nil
	result = SafelyRunPanicIfError(nil, message)
	if result {
		t.Error("should did panic")
	}

	// err is not nil
	result = SafelyRunPanicIfError(errors.New("error"), message)
	if !result {
		t.Error("should did panic")
	}
}

func SafelyRunPanicIfError(err error, message string) bool {
	status := false
	var r interface{}

	func() {
		defer func() {
			if r = recover(); r != nil {
				status = true
			}
		}()
		PanicIfError(err, message)
	}()

	return status
}

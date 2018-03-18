package errutil

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicIfError(t *testing.T) {
	message := "error message"
	err := errors.New("error")

	// err is no nil
	assert.Panics(t, func() { PanicIfError(err, message) })

	// err is nil
	assert.NotPanics(t, func() { PanicIfError(nil, message) })
}

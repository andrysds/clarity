package errutil

import (
	"database/sql"
	"errors"
	"testing"

	_ "github.com/go-sql-driver/mysql"
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

func ExamplePanicIfError() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	// Calls Panic if there is error on connecting mysql.
	PanicIfError(err, "error on connecting mysql")
	return db
}

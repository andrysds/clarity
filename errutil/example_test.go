package errutil_test

import (
	"database/sql"

	"github.com/andrysds/clarity/errutil"
	_ "github.com/go-sql-driver/mysql"
)

func ExamplePanicIfError() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	// Calls Panic if there is error on connecting mysql.
	errutil.PanicIfError(err, "error on connecting mysql")
	return db
}

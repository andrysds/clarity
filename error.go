package clarity

import "fmt"

// PanicIfError calls panic if the given error is not nil
func PanicIfError(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s\n", message, err))
	}
}

// PrintIfError prints the given error if it is not nil
func PrintIfError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s\n", message, err)
	}
}

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"hello-go-errors/lib"
)

func main() {
	lib.PrintSeparater("simple")
	if err := lib.GenerateSimpleError("error dayo"); err != nil {
		fmt.Printf("%+v\n", err)
	}

	lib.PrintSeparater("simple: print時 withstack")
	if err := lib.GenerateSimpleError("error dayo"); err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}

	lib.PrintSeparater("pkg errors")
	if err := lib.GenerateErrorFromPkgError("pkg/error no New"); err != nil {
		fmt.Printf("%+v\n", err)
	}
	lib.PrintSeparater("pkg errors: print時 withstack")
	if err := lib.GenerateErrorFromPkgError("pkg/error no New"); err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}

	lib.PrintSeparater("pkg errors with stack:")
	if err := lib.GenerateErrorWithStack("pkg/error no New + withstack"); err != nil {
		fmt.Printf("%+v\n", err)
	}

	lib.PrintSeparater("pkg errors with stack: print時 withstack")
	if err := lib.GenerateErrorWithStack("pkg/error no New + withstack"); err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}
}

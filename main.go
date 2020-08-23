package main

import (
	"hello-go-errors/lib"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	printSeparater("simple")
	if err := lib.GenerateSimpleError("error dayo"); err != nil {
		fmt.Printf("%+v\n", err)
	}

	printSeparater("simple: print時 withstack")
	if err := lib.GenerateSimpleError("error dayo"); err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}

	printSeparater("pkg errors")
	if err := lib.GenerateErrorFromPkgError("pkg/error no New"); err != nil {
		fmt.Printf("%+v\n", err)
	}
	printSeparater("pkg errors: print時 withstack")
	if err := lib.GenerateErrorFromPkgError("pkg/error no New"); err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}

	printSeparater("pkg errors with stack:")
	if err := lib.GenerateErrorWithStack("pkg/error no New + withstack"); err != nil {
		fmt.Printf("%+v\n", err)
	}

	printSeparater("pkg errors with stack: print時 withstack")
	if err := lib.GenerateErrorWithStack("pkg/error no New + withstack"); err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}
}

func printSeparater(ctx string) {
	fmt.Printf("\n==== %s ====\n", ctx)
}

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
		fmt.Printf("%+v\n", errors.WithStack(err))
	}

	printSeparater("with stack (print)")
	if err := lib.GenerateErrorFromPkgError("pkg/error no New"); err != nil {
		fmt.Printf("%+v\n", err)
	}

	printSeparater("with stack (generate)")
	if err := lib.GenerateErrorWithStack("pkg/error no New + withstack"); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func printSeparater(ctx string) {
	fmt.Printf("\n==== %s ====\n", ctx)
}

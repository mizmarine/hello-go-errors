package main

import (
	"hello-go-errors/lib"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	printSeparater("start")

	if err := lib.GenerateError("error dayo"); err != nil {
		fmt.Printf("%+v", errors.WithStack(err))

		// これはだめ
		//fmt.Println(errors.WithStack(err))
	}
}

func printSeparater(ctx string) {
	fmt.Printf("==== %s ====\n", ctx)
}

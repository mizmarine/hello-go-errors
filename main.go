package main

import (
	"hello-go-errors/lib"
	"fmt"
)

func main() {
	fmt.Println("hello")

	if err := lib.GenerateError("error dayo"); err != nil {
		fmt.Println(err)
	}
}

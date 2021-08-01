package main

import (
	"fizzbuzz/app"
	"fmt"
)

func main() {
	for i := uint8(1); i <= 100; i++ {
		out, err := app.FizzBuzz(i)
		panicIfErr(err)
		fmt.Println(out)
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

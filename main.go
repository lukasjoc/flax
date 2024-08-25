package main

import (
	"fmt"

	"github.com/lukasjoc/flax/flax"
)

func main() {
	program := flax.Shift()

	fmt.Printf("%#v\n", program)

	arg1 := flax.Shift()
	fmt.Printf("%#v\n", arg1)

	arg2 := flax.Shift()
	fmt.Printf("%#v\n", arg2)
}

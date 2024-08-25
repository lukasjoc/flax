package main

import (
	"fmt"

	"github.com/lukasjoc/flax/flax"
)

func main() {
	program := flax.Shift()
	fmt.Printf("%s IsProgram:%v\n", program.Name, program.Program)

    flax.Unshift()
	programAgain := flax.Shift()
	fmt.Printf("%s IsProgram:%v\n", programAgain.Name, programAgain.Program)

	programAgainAgain := flax.Peek()
	if programAgainAgain != nil {
		fmt.Printf("%s IsProgram:%v\n", programAgainAgain.Name, programAgainAgain.Program)
	} else {
		fmt.Println("NIL")
	}
}

package main

import (
	"calc/eval"
	"calc/lexer"
	"fmt"
)

func main() {
	for {
		fmt.Print(" ")

		var input string
		fmt.Scanln(&input)

		val := eval.Evaluate(lexer.NewLexer(input))
		fmt.Printf(" %d\n", val)
	}
}

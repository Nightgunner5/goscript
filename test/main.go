package main

import (
	"fmt"
	"github.com/Nightgunner5/goscript"
	"github.com/Nightgunner5/goscript/parser"
)

func main() {
	var state goscript.State

	input := `{
print
10
-
-
7
;
}`
	fmt.Println("Input: ", input)

	program, err := parser.Parse(input)

	fmt.Println("Errors: ", err)

	fmt.Printf("%+v\n", program)

	program.Execute(&state)
}

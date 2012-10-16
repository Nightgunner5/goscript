package parser_test

import (
	"github.com/Nightgunner5/goscript"
	"github.com/Nightgunner5/goscript/parser"
	"log"
)

func ExampleParse() {
	var state goscript.State

	input := `{
	print 10 - -7;
	print 1 + 4 * 4;
};`

	program, err := parser.Parse(input)

	if err != nil {
		log.Println("Error parsing: ", err)
	}

	program.Execute(&state)

	// Output:
	// 17
	// 17
}

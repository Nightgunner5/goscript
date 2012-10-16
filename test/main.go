package main

import (
	"github.com/Nightgunner5/goscript"
	"github.com/Nightgunner5/goscript/parser"
)

func main() {
	var state goscript.State

	program, _ := parser.Parse(`print 10 + 7;`)

	program.Execute(&state)
}

package main

import "github.com/Nightgunner5/goscript"

func Val(v interface{}) goscript.Value {
	return goscript.Value{
		Value: v,
		Flags: goscript.Flags{},
	}
}

func main() {
	var state goscript.State

	program := goscript.I_block{
		goscript.I_const{
			Value: Val(10.0),
		},
		goscript.I_const{
			Value: Val(7.0),
		},
		goscript.I_math_add,
		goscript.I_print,
	}

	program.Execute(&state)
}

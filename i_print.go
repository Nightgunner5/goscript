package goscript

import "fmt"

type I_print struct {}

func (I_print) Execute(state *State) {
	fmt.Println(state.PopStack().Value)
}

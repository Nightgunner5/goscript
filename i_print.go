package goscript

import "fmt"

type I_print_ byte
const I_print I_print_ = 0

func (I_print_) Execute(state *State) {
	fmt.Println(state.PopStack().Value)
}

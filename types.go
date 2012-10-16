package goscript

type Flag string
type Flags []Flag

type Value struct {
	Value interface{}
	Flags Flags
}

type State struct {
	Parent *State
	Heap   map[string]Value
	Stack  *Stack
}

type Stack struct {
	Parent *Stack
	Value  Value
}

type Instruction interface {
	Execute(*State)
}

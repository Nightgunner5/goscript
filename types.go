package goscript

type Value struct {
	Value interface{}
	Flags Flags
}

type State struct {
	*state
}

type state struct {
	parent *state
	heap   map[string]Value
	stack  *stack
}

type stack struct {
	parent *stack
	value  Value
}

type Instruction interface {
	Execute(*State)
	String() string
}

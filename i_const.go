package goscript

import "fmt"

type I_const struct {
	Value Value
}

func (inst I_const) Execute(state *State) {
	state.PushStack(inst.Value)
}

func (inst I_const) String() string {
	return fmt.Sprintf("const %T %#v %v", inst.Value.Value, inst.Value.Value, inst.Value.Flags)
}

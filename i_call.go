package goscript

import "fmt"

type I_call struct {
	ID string
}

func (inst I_call) Execute(state *State) {
	args := make([]Value, 0, 1)
	for state.state != nil && state.state.stack != nil {
		args = append(args, state.PopStack())
	}
	f := Builtin_Print
	state.PushStack(f(args))
}

func (inst I_call) String() string {
	return fmt.Sprintf("call %q", inst.ID)
}

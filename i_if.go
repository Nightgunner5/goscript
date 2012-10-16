package goscript

import "fmt"

// Stack must have a bool at the top. Path executed is determined by
// the value of the bool.
type I_if struct {
	True, False Instruction
}

func (inst I_if) Execute(state *State) {
	condition := state.PopStack()
	if c, ok := condition.Value.(bool); ok {
		if c {
			inst.True.Execute(state)
		} else {
			inst.False.Execute(state)
		}
	} else {
		state.PushStack(condition)
		panic("Non-bool condition for if statement.")
	}
}

func (inst I_if) String() string {
	return fmt.Sprintf("if\n%s\nelse\n%s", inst.True, inst.False)
}

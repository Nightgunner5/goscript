package goscript

// Stack must have a bool at the top. Path executed is determined by
// the value of the bool.
type I_if struct {
	True, False Instruction
}

func (instruction I_if) Execute(state *State) {
	condition := state.PopStack()
	if c, ok := condition.Value.(bool); ok {
		if c {
			instruction.True.Execute(state)
		} else {
			instruction.False.Execute(state)
		}
	} else {
		state.PushStack(condition)
		panic("Non-bool condition for if statement.")
	}
}

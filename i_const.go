package goscript

type I_const struct {
	Value Value
}

func (instruction I_const) Execute(state *State) {
	state.PushStack(instruction.Value)
}

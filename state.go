package goscript

func (state *State) PushStack(v Value) {
	state.Stack = &Stack{
		Parent: state.Stack,
		Value:  v,
	}
}

func (state *State) PopStack() (v Value) {
	if state.Stack == nil {
		panic("Stack underflow: A nonexistent value was read.")
	}
	v, state.Stack = state.Stack.Value, state.Stack.Parent
	return
}

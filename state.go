package goscript

func (s *State) PushStack(v Value) {
	if s.state == nil {
		s.state = &state{}
	}

	s.state.PushStack(v)
}

func (state *state) PushStack(v Value) {
	state.stack = &stack{
		parent: state.stack,
		value:  v,
	}
}

func (state *state) PopStack() (v Value) {
	if state.stack == nil {
		panic("Stack underflow: A nonexistent value was read.")
	}
	v, state.stack = state.stack.value, state.stack.parent
	return
}

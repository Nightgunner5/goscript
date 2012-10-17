package goscript

type I_state byte

const (
	I_state_push     I_state = '>'
	I_state_pop      I_state = '<'
	I_state_popstack I_state = '.'
)

func (inst I_state) Execute(s *State) {
	switch inst {
	case I_state_push:
		s.state = &state{
			parent: s.state,
			heap:   make(map[string]Value),
			stack:  nil,
		}
	case I_state_pop:
		old := s.state
		s.state = s.state.parent
		s.PushStack(old.PopStack())
	case I_state_popstack:
		s.PopStack()
	default:
		panic("Unknown value for I_state: '" + string([]byte{byte(inst)}) + "'")
	}
}

func (inst I_state) String() string {
	switch inst {
	case I_state_push:
		return "state_push"
	case I_state_pop:
		return "state_pop"
	case I_state_popstack:
		return "state_popstack"
	}
	return "ERROR unknown I_state"
}

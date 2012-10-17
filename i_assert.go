package goscript

type I_assert byte

const (
	I_assert_nostack I_assert = 'S'
)

func (inst I_assert) Execute(state *State) {
	switch inst {
	case I_assert_nostack:
		if state.state != nil && state.state.stack != nil {
			panic("Assertion failed: Stack must be empty.")
		}
	}
}

func (inst I_assert) String() string {
	switch inst {
	case I_assert_nostack:
		return "assert_nostack"
	}
	return "ERROR assert"
}

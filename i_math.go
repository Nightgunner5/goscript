package goscript

// Only works with float64
type I_math byte

const (
	I_math_add I_math = '+'
	I_math_neg I_math = '-'
	I_math_mul I_math = '*'
	I_math_div I_math = '/'
)

func (math I_math) Execute(state *State) {
	switch math {
	case I_math_add, I_math_mul, I_math_div:
		b, a := state.PopStack(), state.PopStack()

		var f func(float64, float64) float64
		switch math {
		case I_math_add:
			f = func(a, b float64) float64 { return a + b }
		case I_math_mul:
			f = func(a, b float64) float64 { return a * b }
		case I_math_div:
			f = func(a, b float64) float64 { return a / b }
		}

		state.PushStack(Value{
			Value: f(a.Value.(float64), b.Value.(float64)),
			Flags: FlagUnion(a.Flags, b.Flags),
		})

	case I_math_neg:
		a := state.PopStack()
		a.Value = -a.Value.(float64)
		state.PushStack(a)
	}
}

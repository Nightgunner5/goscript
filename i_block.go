package goscript

type I_block []Instruction

func (block I_block) Execute(state *State) {
	for _, instruction := range block {
		instruction.Execute(state)
	}
}

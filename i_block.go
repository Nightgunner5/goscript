package goscript

import (
	"runtime"
	"strings"
)

type I_block []Instruction

func (block I_block) Execute(state *State) {
	for _, instruction := range block {
		instruction.Execute(state)
		runtime.Gosched() // Make sure any given script can't take the whole CPU
	}
}

func (block I_block) String() string {
	out := "{"

	for _, inst := range block {
		out += strings.Replace("\n"+inst.String(), "\n", "\n\t", -1)
	}

	return out + "\n}"
}

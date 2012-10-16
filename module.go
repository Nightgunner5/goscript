package goscript

import "sync"

type Module struct {
	Name   string
	Script Instruction
	Uses   []string // Module Names
}

type Domain struct {
	Modules map[string]Module // Key == Module.Name
	Lock    sync.RWMutex
}

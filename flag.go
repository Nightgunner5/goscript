package goscript

type Flag string
type Flags []Flag

func (f Flags) Has(flag Flag) bool {
	for _, have := range f {
		if have == flag {
			return true
		}
	}
	return false
}

func (f *Flags) Set(flag Flag) {
	if !f.Has(flag) {
		*f = append(*f, flag)
	}
}

func (f *Flags) Unset(flag Flag) {
	for i, have := range *f {
		if have == flag {
			*f = append((*f)[:i-1], (*f)[i+1:]...)
			return
		}
	}
}

func FlagUnion(flagsets ...Flags) Flags {
	result := make(Flags, 0)

	for _, flags := range flagsets {
		for _, flag := range flags {
			result.Set(flag)
		}
	}

	return result
}

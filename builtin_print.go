package goscript

import "fmt"

func Builtin_Print(args []Value) Value {
	toPrint := make([]interface{}, 0, len(args))
	for _, arg := range args {
		toPrint = append(toPrint, arg.Value)
	}
	fmt.Println(toPrint...)

	return Value{}
}

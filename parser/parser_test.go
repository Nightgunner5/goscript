package parser_test

import (
	"fmt"
	"github.com/Nightgunner5/goscript"
	"github.com/Nightgunner5/goscript/parser"
	"log"
	"strings"
	"testing"
)

func expectParseError(substring string, script string, t *testing.T) {
	_, err := parser.Parse(script)
	if err == nil {
		t.Errorf("Expected parse error containing %q, but there was no error.", substring)
	} else if !strings.Contains(fmt.Sprint(err), substring) {
		t.Errorf("Expected parse error containing %q, but error was %q.", substring, err)
	}
}

func expectPanic(substring string, script string, t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			if !strings.Contains(fmt.Sprint(p), substring) {
				t.Errorf("Expected panic containing %q, but panic was %q.", substring, p)
			}
		} else {
			t.Errorf("Expected panic containing %q, but no panic.", substring)
		}
	}()

	var state goscript.State

	program, err := parser.Parse(script)
	if err != nil {
		t.Error("Unexpected parse error:", err)
	}

	program.Execute(&state)
}

func TestErrorSemicolon(t *testing.T) {
	expectParseError("last read: \";\"", `A(1 + 1);
	B(2 + 2);;
	C(3 + 3);`, t)
}

func ExampleParse() {
	var state goscript.State

	input := `{
		Print(10 - -7);
		Print(1 + 4 * 4);
	};`

	program, err := parser.Parse(input)

	if err != nil {
		log.Println("Error parsing: ", err)
	}

	program.Execute(&state)

	// Output:
	// 17
	// 17
}

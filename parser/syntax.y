%{
package parser

import (
	"fmt"
	. "github.com/Nightgunner5/goscript"
)
%}

%start program

%union {
	num  float64
	inst []Instruction
}

%type <inst> program block_stmt stmt expr

%token <num> NUMBER
%token print

%left '{' '}'
%left ';'

%left '+' '-'
%left '*' '/'
%left '(' ')'

%%

program:
	program stmt ';'
		{
			$$ = append($1, $2...)
			yylex.(*lexer).inst = $$
		}
|	/* empty */
		{
			$$ = []Instruction{}
		}
;

stmt:
	print expr
		{
			$$ = append($2, I_print)
		}
|	'{' block_stmt '}'
		{
			$$ = []Instruction{
				I_block($2),
			}
		}
;

block_stmt:
	block_stmt stmt ';'
		{
			$$ = append($1, $2...)
		}
|	/* empty */
		{
			$$ = []Instruction{}
		}

expr:
	'(' expr ')'
		{
			$$ = $2
		}
|	expr '+' expr
		{
			$$ = append(append($1, $3...), I_math_add)
		}
|	expr '-' expr
		{
			$$ = append(append($1, $3...), I_math_neg, I_math_add)
		}
|	'-' expr
		{
			$$ = append($2, I_math_neg)
		}
|	expr '*' expr
		{
			$$ = append(append($1, $3...), I_math_mul)
		}
|	expr '/' expr
		{
			$$ = append(append($1, $3...), I_math_div)
		}
|	NUMBER
		{
			$$ = []Instruction{
				I_const{
					Value: Value{
						Value: $1,
						Flags: Flags{},
					},
				},
			}
		}
;

%%

type multiError []string

func (err multiError) Error() string {
	// If there's a function that does this more efficiently,
	// I'm too lazy to look for it.
	out := err[0]

	for i := range err {
		if i == 0 {
			continue
		}

		out += "\n" + err[i]
	}

	return out
}

type lexer struct {
	source []rune
	err    multiError
	line   int
	inst   []Instruction
}

func (lex *lexer) hasText(text string) bool {
	t := []rune(text)

	if len(t) > len(lex.source) {
		return false
	}

	for i, c := range t {
		if lex.source[i] != c {
			return false
		}
	}

	lex.source = lex.source[len(t):]
	return true
}

func (lex *lexer) Lex(lval *yySymType) int {
	c := ' '

	for c == ' ' {
		if len(lex.source) == 0 {
			return 0
		}
		c = lex.source[0]
		lex.source = lex.source[1:]
	}

	if c == '\n' {
		lex.line++
		return lex.Lex(lval)
	}

	if c == '(' || c == ')' || c == '+' || c == '-' || c == '*' || c == '/' || c == ';' || c == '{' || c == '}' {
		return int(c)
	}

	if c >= '0' && c <= '9' {
		lval.num = float64(c - '0')

		for lex.source[0] >= '0' && lex.source[0] <= '9' {
			lval.num *= 10
			lval.num += float64(lex.source[0] - '0')
			lex.source = lex.source[1:]
		}

		return NUMBER
	}

	if c == 'p' && lex.hasText("rint") {
		return print
	}

	return yyErrCode
}

func (lex *lexer) Error(e string) {
	lex.err = append(lex.err, fmt.Sprintf("%s (on line %d)", e, lex.line))
}

func Parse(source string) (Instruction, error) {
	l := &lexer{
		source: []rune(source),
		err: nil,
		line: 1,
	}

	yyParse(l)

	if len(l.err) == 0 {
		return I_block(l.inst), nil
	}
	return I_block(l.inst), l.err
}

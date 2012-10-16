%{
package parser

import . "github.com/Nightgunner5/goscript"
%}

%start program

%union {
	num  float64
	inst []Instruction
}

%type <inst> program stmt expr

%token <num> NUMBER
%token print

%left ';'

%left '+'
%left '(' ')'

%%

program:
	program stmt ';'
		{
			$$ = append($1, $2...)
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
;

expr:
	'(' expr ')'
		{
			$$ = $2
		}
|	expr '+' expr
		{
			$$ = append(append($1, $3...), I_math_add)
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
	first   *yySymType
}

func (lex *lexer) Lex(lval *yySymType) int {
	if lex.first == nil {
		lex.first = lval
	}

	c := ' '

	for c == ' ' {
		if len(lex.source) == 0 {
			return yyEofCode
		}
		c = lex.source[0]
		lex.source = lex.source[1:]
	}

	if c == '(' || c == ')' || c == '+' || c == ';' {
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

	if c == 'p' && lex.source[0] == 'r' && lex.source[1] == 'i' && lex.source[2] == 'n' && lex.source[3] == 't' && lex.source[4] == ' ' {
		lex.source = lex.source[4:]
		return print
	}

	__yyfmt__.Println(string([]rune{c}) + string(lex.source))

	return yyErrCode
}

func (lex *lexer) Error(e string) {
	lex.err = append(lex.err, e)
}

func Parse(source string) ([]Instruction, error) {
	yyDebug = 1000

	l := &lexer{
		source: []rune(source),
		err: nil,
	}

	yyParse(l)

	__yyfmt__.Printf("%#v\n", l.first)

	return l.first.inst, l.err
}

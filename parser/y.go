//line syntax.y:2
package parser

import __yyfmt__ "fmt"

//line syntax.y:2
import (
	"fmt"
	. "github.com/Nightgunner5/goscript"
	"unicode"
)

//line syntax.y:13
type yySymType struct {
	yys   int
	num   float64
	ident string
	inst  []Instruction
}

const identifier = 57346
const number = 57347

var yyToknames = []string{
	"identifier",
	"number",
	" {",
	" }",
	" ;",
	" +",
	" -",
	" *",
	" /",
	" (",
	" )",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line syntax.y:125

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
	read   []rune
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

	lex.read = append(lex.read, t...)
	lex.source = lex.source[len(t):]
	return true
}

func (lex *lexer) Lex(lval *yySymType) int {
	lex.read = nil
	c := ' '

	for c == ' ' || c == '\t' {
		if len(lex.source) == 0 {
			return 0
		}
		c = lex.source[0]
		lex.read = append(lex.read, c)
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
			lex.read = append(lex.read, lex.source[0])
			lex.source = lex.source[1:]
		}

		return number
	}

	if unicode.IsLetter(c) || c == '_' {
		var i int
		for i = 0; i < len(lex.source); i++ {
			if !unicode.IsLetter(lex.source[i]) &&
				!unicode.IsDigit(lex.source[i]) &&
				lex.source[i] != '_' {
				break
			}
		}
		lval.ident = string(append([]rune{c}, lex.source[:i]...))
		lex.read = append(lex.read, lex.source[:i]...)
		lex.source = lex.source[i:]
		return identifier
	}

	return yyErrCode
}

func (lex *lexer) Error(e string) {
	lastRead := string(lex.read)
	if lastRead == "" {
		lex.err = append(lex.err, fmt.Sprintf("%s (on line %d) (at EOF)", e, lex.line))
	} else {
		lex.err = append(lex.err, fmt.Sprintf("%s (on line %d) (last read: %q)", e, lex.line, lastRead))
	}
}

func Parse(source string) (Instruction, error) {
	l := &lexer{
		source: []rune(source),
		err:    nil,
		line:   1,
	}

	yyParse(l)

	if len(l.err) == 0 {
		return I_block(l.inst), nil
	}
	return I_block(l.inst), l.err
}

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 17
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 44

var yyAct = []int{

	9, 8, 18, 19, 20, 21, 25, 31, 33, 17,
	24, 22, 23, 16, 17, 20, 21, 5, 26, 27,
	28, 29, 30, 12, 13, 6, 32, 7, 1, 11,
	2, 0, 10, 18, 19, 20, 21, 3, 15, 4,
	14, 3, 0, 4,
}
var yyPact = []int{

	-1000, 37, 9, 12, -1000, -1000, 19, 33, -1, 24,
	19, 19, -3, -1000, -1000, -2, -1000, 19, 19, 19,
	19, 19, -7, 4, 19, -1000, 24, 4, 4, -1000,
	-1000, -1000, -6, -1000,
}
var yyPgo = []int{

	0, 28, 27, 30, 0, 1,
}
var yyR1 = []int{

	0, 1, 1, 3, 3, 2, 2, 4, 4, 4,
	4, 4, 4, 4, 4, 5, 5,
}
var yyR2 = []int{

	0, 3, 0, 4, 3, 3, 0, 3, 3, 3,
	2, 3, 3, 4, 1, 3, 1,
}
var yyChk = []int{

	-1000, -1, -3, 4, 6, 8, 13, -2, -5, -4,
	13, 10, 4, 5, 7, -3, 14, 15, 9, 10,
	11, 12, -4, -4, 13, 8, -4, -4, -4, -4,
	-4, 14, -5, 14,
}
var yyDef = []int{

	2, -2, 0, 0, 6, 1, 0, 0, 0, 16,
	0, 0, 0, 14, 4, 0, 3, 0, 0, 0,
	0, 0, 0, 10, 0, 5, 15, 8, 9, 11,
	12, 7, 0, 13,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	13, 14, 11, 9, 15, 10, 3, 12, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 8,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 6, 3, 7,
}
var yyTok2 = []int{

	2, 3, 4, 5,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c > 0 && c <= len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line syntax.y:35
		{
			yyVAL.inst = append(yyS[yypt-2].inst, yyS[yypt-1].inst...)
			yylex.(*lexer).inst = yyVAL.inst
		}
	case 2:
		//line syntax.y:40
		{
			yyVAL.inst = []Instruction{}
		}
	case 3:
		//line syntax.y:47
		{
			yyVAL.inst = append(yyS[yypt-1].inst, I_call{
				ID: yyS[yypt-3].ident,
			}, I_state_popstack)
		}
	case 4:
		//line syntax.y:53
		{
			yyVAL.inst = []Instruction{
				I_block(yyS[yypt-1].inst),
			}
		}
	case 5:
		//line syntax.y:62
		{
			yyVAL.inst = append(yyS[yypt-2].inst, yyS[yypt-1].inst...)
		}
	case 6:
		//line syntax.y:66
		{
			yyVAL.inst = []Instruction{}
		}
	case 7:
		//line syntax.y:72
		{
			yyVAL.inst = yyS[yypt-1].inst
		}
	case 8:
		//line syntax.y:76
		{
			yyVAL.inst = append(append(yyS[yypt-2].inst, yyS[yypt-0].inst...), I_math_add)
		}
	case 9:
		//line syntax.y:80
		{
			yyVAL.inst = append(append(yyS[yypt-2].inst, yyS[yypt-0].inst...), I_math_neg, I_math_add)
		}
	case 10:
		//line syntax.y:84
		{
			yyVAL.inst = append(yyS[yypt-0].inst, I_math_neg)
		}
	case 11:
		//line syntax.y:88
		{
			yyVAL.inst = append(append(yyS[yypt-2].inst, yyS[yypt-0].inst...), I_math_mul)
		}
	case 12:
		//line syntax.y:92
		{
			yyVAL.inst = append(append(yyS[yypt-2].inst, yyS[yypt-0].inst...), I_math_div)
		}
	case 13:
		//line syntax.y:96
		{
			yyVAL.inst = append(append([]Instruction{I_state_push}, yyS[yypt-1].inst...), I_call{
				ID: yyS[yypt-3].ident,
			}, I_state_pop)
		}
	case 14:
		//line syntax.y:102
		{
			yyVAL.inst = []Instruction{
				I_const{
					Value: Value{
						Value: yyS[yypt-0].num,
						Flags: Flags{},
					},
				},
			}
		}
	case 15:
		//line syntax.y:116
		{
			yyVAL.inst = append(yyS[yypt-2].inst, yyS[yypt-0].inst...)
		}
	case 16:
		//line syntax.y:120
		{
			yyVAL.inst = yyS[yypt-0].inst
		}
	}
	goto yystack /* stack new state and value */
}

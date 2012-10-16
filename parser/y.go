
//line syntax.y:2
package parser
import __yyfmt__ "fmt"
//line syntax.y:2
		
import . "github.com/Nightgunner5/goscript"

//line syntax.y:9
type yySymType struct {
	yys int
	num  float64
	inst []Instruction
}

const NUMBER = 57346
const print = 57347

var yyToknames = []string{
	"NUMBER",
	"print",
	" ;",
	" +",
	" (",
	" )",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line syntax.y:66


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

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 7
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 12

var yyAct = []int{

	8, 5, 11, 7, 8, 2, 4, 6, 9, 3,
	10, 1,
}
var yyPact = []int{

	-1000, 4, 0, -1, -1000, -3, -1, -1000, -1, -7,
	-1000, -1000,
}
var yyPgo = []int{

	0, 11, 5, 1,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 3, 3,
}
var yyR2 = []int{

	0, 3, 0, 2, 3, 3, 1,
}
var yyChk = []int{

	-1000, -1, -2, 5, 6, -3, 8, 4, 7, -3,
	-3, 9,
}
var yyDef = []int{

	2, -2, 0, 0, 1, 3, 0, 6, 0, 0,
	5, 4,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	8, 9, 3, 7, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 6,
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
		//line syntax.y:28
		{
				yyVAL.inst = append(yyS[yypt-2].inst, yyS[yypt-1].inst...)
			}
	case 2:
		//line syntax.y:32
		{
				yyVAL.inst = []Instruction{}
			}
	case 3:
		//line syntax.y:39
		{
				yyVAL.inst = append(yyS[yypt-0].inst, I_print)
			}
	case 4:
		//line syntax.y:46
		{
				yyVAL.inst = yyS[yypt-1].inst
			}
	case 5:
		//line syntax.y:50
		{
				yyVAL.inst = append(append(yyS[yypt-2].inst, yyS[yypt-0].inst...), I_math_add)
			}
	case 6:
		//line syntax.y:54
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
	}
	goto yystack /* stack new state and value */
}

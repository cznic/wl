// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"fmt"
	"go/token"
	"io"
	"strings"
	"unicode"

	"github.com/cznic/golex/lex"
)

const (
	ccEOF = iota + 0x80
	ccDigit
	ccLetter
	ccLetterLike
	ccIgnore
	ccConjugate
	ccTranspose
	ccConjugateTranspose
	ccHermitianConjugate
	ccSqrt
	ccIntegrate
	ccDifferentialD
	ccPartialD
	ccDel
	ccDiscreteShift
	ccDiscreteRatio
	ccDifferenceDelta
	ccSquare
	ccSmallCircle
	ccCircleDot
	ccCross
	ccPlusMinus
	ccMinusPlus
	ccDivide
	ccBackslash

	ccOther
)

var (
	letterLike = &unicode.RangeTable{
		R16: []unicode.Range16{
			{'\u2100', '\u214f', 1},
		},
	}

	namedChars = map[string]rune{
		"Backslash":           '\u2216',
		"CircleDot":           '\u2299',
		"Conjugate":           '\uf3c8',
		"ConjugateTranspose":  '\uf3c9',
		"Cross":               '\uf4a0',
		"Del":                 '\u2207',
		"DifferenceDelta":     '\u2206',
		"DifferentialD":       '\uf74c',
		"DiscreteRatio":       '\uf4a4',
		"DiscreteShift":       '\uf4a3',
		"Divide":              '\u00f7',
		"HermitianConjugate":  '\uf3ce',
		"Integrate":           '\u222b',
		"MinusPlus":           '\u2213',
		"PartialD":            '\u2202',
		"PlusMinus":           '\u00b1',
		"RawAmpersand":        '&',
		"RawAt":               '@',
		"RawBackquote":        '`',
		"RawBackslash":        '∖',
		"RawColon":            ':',
		"RawComma":            ',',
		"RawDash":             '-',
		"RawDollar":           '$',
		"RawDot":              '.',
		"RawDoubleQuote":      '"',
		"RawEqual":            '=',
		"RawExclamation":      '!',
		"RawGreater":          '>',
		"RawLeftBrace":        '{',
		"RawLeftBracket":      '[',
		"RawLeftParenthesis":  '(',
		"RawLess":             '<',
		"RawNumberSign":       '#',
		"RawPercent":          '%',
		"RawPlus":             '+',
		"RawQuestion":         '?',
		"RawQuote":            '\'',
		"RawRightBrace":       '}',
		"RawRightBracket":     ']',
		"RawRightParenthesis": ')',
		"RawSemicolon":        ';',
		"RawSlash":            '/',
		"RawSpace":            ' ',
		"RawStar":             '*',
		"RawTilde":            '~',
		"RawUnderscore":       '_',
		"RawVerticalBar":      '|',
		"RawWedge":            '^',
		"SmallCircle":         '\u2218',
		"Sqrt":                '\u221a',
		"Square":              '\uf520',
		"Transpose":           '\uf3c7',
	}
)

func runeClass(r rune) int {
	switch {
	case r == lex.RuneEOF:
		return ccEOF
	case r == IGNORE:
		return ccIgnore
	case r < 0x80:
		return int(r)
	case r == '\uf3c7':
		return ccTranspose
	case r == '\uf3c8':
		return ccConjugate
	case r == '\uf3c9':
		return ccConjugateTranspose
	case r == '\uf3ce':
		return ccHermitianConjugate
	case r == '\u221a':
		return ccSqrt
	case r == '\u222b':
		return ccIntegrate
	case r == '\uf74c':
		return ccDifferentialD
	case r == '\u2202':
		return ccPartialD
	case r == '\u2207':
		return ccDel
	case r == '\uf4a3':
		return ccDiscreteShift
	case r == '\uf4a4':
		return ccDiscreteRatio
	case r == '\u2206':
		return ccDifferenceDelta
	case r == '\uf520':
		return ccSquare
	case r == '\u2218':
		return ccSmallCircle
	case r == '\u2299':
		return ccCircleDot
	case r == '\uf4a0':
		return ccCross
	case r == '\u00b1':
		return ccPlusMinus
	case r == '\u2213':
		return ccMinusPlus
	case r == '\u00f7':
		return ccDivide
	case r == '\u2216':
		return ccBackslash
	case unicode.IsDigit(r):
		return ccDigit
	case unicode.IsLetter(r):
		return ccLetter
	case unicode.Is(letterLike, r):
		return ccLetterLike
	default:
		return ccOther
	}
}

type lexer struct {
	*lex.Lexer
	ast         Node
	c           rune
	err         error
	exampleAST  interface{}
	exampleRule int
	in          []rune
	interactive bool
	mark        int
	r           io.RuneReader
	rerr        error
	sc          int // Start condition.
	stack       []int
	str         []byte
	sz          int
	un          []rune
}

func newLexer(r io.RuneReader) *lexer {
	return &lexer{
		c:           -1,
		exampleRule: -1,
		r:           r,
	}
}

func (lx *lexer) init(l *lex.Lexer, interactive bool) {
	lx.Lexer = l
	lx.ast = nil
	lx.c = -1
	lx.err = nil
	lx.exampleAST = nil
	lx.in = lx.in[:0]
	lx.interactive = interactive
	lx.mark = -1
	lx.sc = 0
	lx.stack = lx.stack[:0]
	lx.str = lx.str[:0]
	lx.un = lx.un[:0]
}

func (lx *lexer) unget(r rune) {
	lx.un = append(lx.un, r)
}

func (lx *lexer) next() (r int) {
	// fmt.Printf("%T.next\n", lx)
	// defer func() { fmt.Printf("\t%T.next: %U\n", lx, r) }()
	if len(lx.un) != 0 {
		r = int(lx.un[len(lx.un)-1])
		lx.un = lx.un[:len(lx.un)-1]
		lx.c = rune(r)
		return r
	}

	lx.in = append(lx.in, lx.c)
	// fmt.Printf("%T.r.ReadRune\n", lx)
	lx.c, lx.sz, lx.rerr = lx.r.ReadRune()
	// fmt.Printf("\t%T.r.ReadRune: %U\n", lx, lx.c)
	if lx.rerr != nil {
		lx.c = -1
		return -1
	}

	return int(lx.c)
}

func (lx *lexer) named(name string) (rune, int, error) {
	c, ok := namedChars[name]
	if !ok {
		return 0, 1, fmt.Errorf("unknown character name %q", name)
	}

	return c, len(name) + len("\\[]"), nil
}

func (lx *lexer) token() string { return string(lx.in) }

func (lx *lexer) position(pos token.Pos) token.Position { return lx.File.Position(pos) }
func (lx *lexer) push(r int) int                        { lx.stack = append(lx.stack, r); return r }

func (lx *lexer) pop() (r int) {
	if n := len(lx.stack); n > 0 {
		r = lx.stack[n-1]
		lx.stack = lx.stack[:n-1]
	}
	return r
}

// func (lx *lexer) unget(r rune) int {
// }

func (lx *lexer) sdump() string {
	var a []string
	for _, v := range lx.stack {
		a = append(a, yySymName(v))
	}
	return fmt.Sprintf("[%v]", strings.Join(a, ", "))
}

func (lx *lexer) errPos(pos token.Pos, msg string) {
	if lx.err == nil {
		lx.err = fmt.Errorf("%s: %v", lx.position(pos), msg)
	}
}

// Implements yyLexer.
func (lx *lexer) Error(msg string) {
	msg = strings.Replace(msg, "$end", "EOF", -1)
	lx.errPos(lx.First.Pos(), msg)
}

// Implements yyLexer.
func (lx *lexer) Lex(lval *yySymType) (r int) {
	// fmt.Printf("%T.Lex\n", lx)
	// defer func() { fmt.Printf("\t%T.Lex: %U %s\n", lx, r, yySymName(r)) }()
	if lx.err != nil {
		// fmt.Printf("\t%T.Lex err %q", lx, lx.err)
		return -1
	}
more:
	r = lx.scan()
	if r == IGNORE {
		// fmt.Printf("\t%T.Lex ignore", lx)
		goto more
	}

	if r == '\n' {
		if lx.interactive && len(lx.stack) == 0 {
			for _, sym := range yyFollow[lval.yys] {
				if sym == yyEofCode {
					return -1
				}
			}
		}

		goto more
	}

	tok := Token{Rune: rune(r), pos: lx.First.Pos()}
	if r > 0x7f {
		switch r {
		case STRING:
			tok.Val = string(lx.str)
		default:
			tok.Val = string(lx.TokenBytes(nil))
		}
	}
	lval.Token = tok
	return r
}

// Implements yyLexerEx.
func (lx *lexer) Reduced(rule, state int, lval *yySymType) (stop bool) {
	if rule == 1 {
		lx.ast = lval.Node
	}
	if rule != lx.exampleRule {
		return false
	}

	switch x := lval.Node.(type) {
	case interface {
		fragment() interface{}
	}:
		lx.exampleAST = x.fragment()
	default:
		lx.exampleAST = x
	}
	return true
}

func (lx *lexer) parse(l *lex.Lexer, interactive bool) error {
	lx.init(l, interactive)
	if yyParse(lx) != 0 && lx.err == nil {
		return fmt.Errorf("%T.parse: internal error", lx)
	}

	return lx.err
}

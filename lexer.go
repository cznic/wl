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
	ccDiamond
	ccWedge
	ccVee
	ccCircleTimes
	ccCenterDot
	ccTimes
	ccStar
	ccProduct
	ccVerticalTilde
	ccCoproduct
	ccCap
	ccCup
	ccCirclePlus
	ccCircleMinus
	ccSum
	ccIntersection
	ccUnion
	ccEqual
	ccVerticalBar
	ccNotVerticalBar
	ccDoubleVerticalBar
	ccNotDoubleVerticalBar
	ccElement
	ccNotElement
	ccSubset
	ccSuperset
	ccForAll
	ccExists
	ccNotExists
	ccNot
	ccAnd
	ccNand
	ccXor
	ccXnor
	ccOr
	ccNor
	ccEquivalent
	ccImplies
	ccRightTee
	ccDoubleRightTee
	ccLeftTee
	ccDoubleLeftTee
	ccUpTee
	ccDownTee
	ccSuchThat
	ccVerticalSeparator
	ccTherefore
	ccBecause
	ccFunction

	ccOther
)

var (
	letterLike = &unicode.RangeTable{
		R16: []unicode.Range16{
			{'\u2100', '\u214f', 1},
		},
	}

	namedChars = map[string]rune{
		"And":                '\u2227',
		"Backslash":          '\u2216',
		"Because":            '\u2235',
		"Cap":                '\u2322',
		"CenterDot":          '\u00b7',
		"CircleDot":          '\u2299',
		"CircleMinus":        '\u2296',
		"CirclePlus":         '\u2295',
		"CircleTimes":        '\u2297',
		"Conjugate":          '\uf3c8',
		"ConjugateTranspose": '\uf3c9',
		"Coproduct":          '\u2210',
		"Cross":              '\uf4a0',
		"Cup":                '\u2323',
		"Del":                '\u2207',
		"Diamond":            '\u22c4',
		"DifferenceDelta":    '\u2206',
		"DifferentialD":      '\uf74c',
		"DiscreteRatio":      '\uf4a4',
		"DiscreteShift":      '\uf4a3',
		"Divide":             '\u00f7',
		"DoubleLeftTee":      '\u2ae4',
		"DoubleRightTee":     '\u22a8',
		"DoubleVerticalBar":  '\u2225',
		"DownTee":            '\u22a4',
		"Element":            '\u2208',
		"Equal":              '\uf431',
		"Equivalent":         '\u29e6',
		"Exists":             '\u2203',
		"ForAll":             '\u2200',
		"Function":           '\uf4a1',
		"HermitianConjugate": '\uf3ce',
		"Implies":            '\uf523',
		"Integrate":          '\u222b',
		"Intersection":       '\u22c2',
		"LeftTee":            '\u22a3',
		"LongEqual":          '\uf7d9',
		"MinusPlus":          '\u2213',
		"Nand":               '\u22bc',
		"Nor":                '\u22bd',
		"Not":                '\u00ac',
		"NotDoubleVerticalBar": '\u2226',
		"NotElement":           '\u2209',
		"NotExists":            '\u2204',
		"NotVerticalBar":       '\uf3d1',
		"Or":                   '\u2228',
		"PartialD":             '\u2202',
		"PlusMinus":            '\u00b1',
		"Product":              '\u220f',
		"RawAmpersand":         '&',
		"RawAt":                '@',
		"RawBackquote":         '`',
		"RawBackslash":         '∖',
		"RawColon":             ':',
		"RawComma":             ',',
		"RawDash":              '-',
		"RawDollar":            '$',
		"RawDot":               '.',
		"RawDoubleQuote":       '"',
		"RawEqual":             '=',
		"RawExclamation":       '!',
		"RawGreater":           '>',
		"RawLeftBrace":         '{',
		"RawLeftBracket":       '[',
		"RawLeftParenthesis":   '(',
		"RawLess":              '<',
		"RawNumberSign":        '#',
		"RawPercent":           '%',
		"RawPlus":              '+',
		"RawQuestion":          '?',
		"RawQuote":             '\'',
		"RawRightBrace":        '}',
		"RawRightBracket":      ']',
		"RawRightParenthesis":  ')',
		"RawSemicolon":         ';',
		"RawSlash":             '/',
		"RawSpace":             ' ',
		"RawStar":              '*',
		"RawTilde":             '~',
		"RawUnderscore":        '_',
		"RawVerticalBar":       '|',
		"RawWedge":             '^',
		"RightTee":             '\u22a2',
		"SmallCircle":          '\u2218',
		"Sqrt":                 '\u221a',
		"Square":               '\uf520',
		"Star":                 '\u22c6',
		"Subset":               '\u2282',
		"SuchThat":             '\u220d',
		"Sum":                  '\u2211',
		"Superset":             '\u2283',
		"Therefore":            '\u2234',
		"Transpose":            '\uf3c7',
		"Union":                '\u22c3',
		"UpTee":                '\u22a5',
		"Vee":                  '\u22c1',
		"VerticalBar":          '\uf3d0',
		"VerticalSeparator":    '\uf432',
		"VerticalTilde":        '\u2240',
		"Wedge":                '\u22c0',
		"Xnor":                 '\uf4a2',
		"Xor":                  '\u22bb',
	}

	classes = map[rune]int{
		'\u00ac': ccNot,
		'\u00b1': ccPlusMinus,
		'\u00b7': ccCenterDot,
		'\u00d7': ccTimes,
		'\u00f7': ccDivide,
		'\u2200': ccForAll,
		'\u2202': ccPartialD,
		'\u2203': ccExists,
		'\u2204': ccNotExists,
		'\u2206': ccDifferenceDelta,
		'\u2207': ccDel,
		'\u2208': ccElement,
		'\u2209': ccNotElement,
		'\u220d': ccSuchThat,
		'\u220f': ccProduct,
		'\u2210': ccCoproduct,
		'\u2211': ccSum,
		'\u2213': ccMinusPlus,
		'\u2216': ccBackslash,
		'\u2218': ccSmallCircle,
		'\u221a': ccSqrt,
		'\u2225': ccDoubleVerticalBar,
		'\u2226': ccNotDoubleVerticalBar,
		'\u2227': ccAnd,
		'\u2228': ccOr,
		'\u222b': ccIntegrate,
		'\u2234': ccTherefore,
		'\u2235': ccBecause,
		'\u2240': ccVerticalTilde,
		'\u2282': ccSubset,
		'\u2283': ccSuperset,
		'\u2295': ccCirclePlus,
		'\u2296': ccCircleMinus,
		'\u2297': ccCircleTimes,
		'\u2299': ccCircleDot,
		'\u22a2': ccRightTee,
		'\u22a3': ccLeftTee,
		'\u22a4': ccDownTee,
		'\u22a5': ccUpTee,
		'\u22a8': ccDoubleRightTee,
		'\u22bc': ccNand,
		'\u22bd': ccNor,
		'\u22bb': ccXor,
		'\u22c0': ccWedge,
		'\u22c1': ccVee,
		'\u22c2': ccIntersection,
		'\u22c3': ccUnion,
		'\u22c4': ccDiamond,
		'\u22c6': ccStar,
		'\u2322': ccCap,
		'\u2323': ccCup,
		'\u29e6': ccEquivalent,
		'\u2ae4': ccDoubleLeftTee,
		'\uf3c7': ccTranspose,
		'\uf3c8': ccConjugate,
		'\uf3c9': ccConjugateTranspose,
		'\uf3ce': ccHermitianConjugate,
		'\uf3d0': ccVerticalBar,
		'\uf3d1': ccNotVerticalBar,
		'\uf431': ccEqual,
		'\uf432': ccVerticalSeparator,
		'\uf4a0': ccCross,
		'\uf4a1': ccFunction,
		'\uf4a2': ccXnor,
		'\uf4a3': ccDiscreteShift,
		'\uf4a4': ccDiscreteRatio,
		'\uf520': ccSquare,
		'\uf523': ccImplies,
		'\uf74c': ccDifferentialD,
		'\uf7d9': ccEqual,
	}
)

func runeClass(r rune) int {
	if cc, ok := classes[r]; ok {
		return cc
	}

	switch {
	case r == lex.RuneEOF:
		return ccEOF
	case r == IGNORE:
		return ccIgnore
	case r < 0x80:
		return int(r)
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
	ast          Node
	c            rune
	commentLevel int
	err          error
	exampleAST   interface{}
	exampleRule  int
	in           []rune
	interactive  bool
	mark         int
	r            io.RuneReader
	rerr         error
	sc           int // Start condition.
	stack        []int
	str          []byte
	sz           int
	un           []rune
	ungetTok     Token
}

func newLexer(r io.RuneReader) *lexer {
	return &lexer{
		c:           -1,
		exampleRule: -1,
		r:           r,
		ungetTok:    Token{Rune: -1},
	}
}

func (lx *lexer) init(l *lex.Lexer, interactive bool) {
	lx.Lexer = l
	lx.ast = nil
	lx.c = -1
	lx.commentLevel = 0
	lx.err = nil
	lx.exampleAST = nil
	lx.in = lx.in[:0]
	lx.interactive = interactive
	lx.mark = -1
	lx.sc = 0
	lx.stack = lx.stack[:0]
	lx.str = lx.str[:0]
	lx.un = lx.un[:0]
	lx.ungetTok.Rune = -1
}

func (lx *lexer) unget(r rune) {
	lx.un = append(lx.un, r)
}

func (lx *lexer) next() (r int) {
	if len(lx.un) != 0 {
		r = int(lx.un[len(lx.un)-1])
		lx.un = lx.un[:len(lx.un)-1]
		lx.c = rune(r)
		return r
	}

	lx.in = append(lx.in, lx.c)
	lx.c, lx.sz, lx.rerr = lx.r.ReadRune()
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
	if lx.err != nil {
		return -1
	}

	if tok := lx.ungetTok; tok.Rune >= 0 {
		lval.Token = tok
		lx.ungetTok.Rune = -1
		return int(tok.Rune)
	}

more:
	r = lx.scan()
	if r == IGNORE {
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

	var ok, mul bool
out:
	for _, sym := range yyFollow[lval.yys] {
		switch sym {
		case r:
			ok = true
			break out
		case '*':
			mul = r >= 0
		}
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
	if !ok && mul {
		lx.ungetTok = tok
		r = '*'
		tok.Rune = rune(r)
		tok.Val = ""
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

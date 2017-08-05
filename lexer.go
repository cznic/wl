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
	ccOther
)

var (
	letterLike = &unicode.RangeTable{
		R16: []unicode.Range16{
			{'\u2100', '\u214f', 1},
		},
	}

	namedChars = map[string]rune{
		"RawSpace":            ' ',
		"RawExclamation":      '!',
		"RawDoubleQuote":      '"',
		"RawNumberSign":       '#',
		"RawDollar":           '$',
		"RawPercent":          '%',
		"RawAmpersand":        '&',
		"RawQuote":            '\'',
		"RawLeftParenthesis":  '(',
		"RawRightParenthesis": ')',
		"RawStar":             '*',
		"RawPlus":             '+',
		"RawComma":            ',',
		"RawDash":             '-',
		"RawDot":              '.',
		"RawSlash":            '/',
		"RawColon":            ':',
		"RawSemicolon":        ';',
		"RawLess":             '<',
		"RawEqual":            '=',
		"RawGreater":          '>',
		"RawQuestion":         '?',
		"RawAt":               '@',
		"RawLeftBracket":      '[',
		"RawBackslash":        '∖',
		"RawRightBracket":     ']',
		"RawWedge":            '^',
		"RawUnderscore":       '_',
		"RawBackquote":        '`',
		"RawLeftBrace":        '{',
		"RawVerticalBar":      '|',
		"RawRightBrace":       '}',
		"RawTilde":            '~',
	}
)

func runeClass(r rune) int {
	switch {
	case r == lex.RuneEOF:
		return ccEOF
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
	ast         Node
	buf         []rune
	c           rune
	err         error
	exampleAST  interface{}
	exampleRule int
	interactive bool
	r           io.RuneReader
	rerr        error
	stack       []int
	sz          int
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
	lx.interactive = interactive
	lx.stack = lx.stack[:0]
}

func (lx *lexer) enter() (r int) {
	lx.buf = lx.buf[:0]
	if lx.c < 0 {
		return lx.next()
	}

	lx.buf = append(lx.buf, lx.c)
	return int(lx.c)
}

func (lx *lexer) next() (r int) {
	lx.c = -1
	lx.c, lx.sz, lx.rerr = lx.r.ReadRune()
	lx.buf = append(lx.buf, lx.c)
	return int(lx.c)
}

func (lx *lexer) named(name string) (rune, int, error) {
	c, ok := namedChars[name]
	if !ok {
		return 0, 1, fmt.Errorf("unknown character name %q", name)
	}

	return c, len(name) + len("\\[]"), nil
}

func (lx *lexer) token() string { return string(lx.buf) }

func (lx *lexer) position(pos token.Pos) token.Position { return lx.File.Position(pos) }
func (lx *lexer) push(r int) int                        { lx.stack = append(lx.stack, r); return r }

func (lx *lexer) pop() (r int) {
	if n := len(lx.stack); n > 0 {
		r = lx.stack[n-1]
		lx.stack = lx.stack[:n-1]
	}
	return r
}

func (lx *lexer) unget(r rune) int {
	if la := lx.Lookahead(); la.Rune != 0 {
		lx.Unget(la)
	}
	lx.Unget(lex.NewChar(lx.First.Pos()+1, r))
	return int(r)
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
more:
	r = lx.scan()
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
		tok.Val = string(lx.TokenBytes(nil))
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

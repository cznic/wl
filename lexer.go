// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"fmt"
	"go/token"
	"strings"
	"unicode"

	"github.com/cznic/golex/lex"
)

const (
	ccEOF = iota + 0x80
	ccDigit
	ccLetter
	ccOther
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
	default:
		return ccOther
	}
}

type lexer struct {
	*lex.Lexer
	ast         Node
	err         error
	exampleAST  interface{}
	exampleRule int
	interactive bool
	stack       []int
}

func newLexer() *lexer { return &lexer{exampleRule: -1} }

func (lx *lexer) init(l *lex.Lexer, interactive bool) {
	lx.Lexer = l
	lx.ast = nil
	lx.err = nil
	lx.exampleAST = nil
	lx.interactive = interactive
	lx.stack = lx.stack[:0]
}

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

	lval.Token = Token{Rune: rune(r), Val: string(lx.TokenBytes(nil)), pos: lx.First.Pos()}
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

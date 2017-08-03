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
	nest        int
}

func newLexer(lx *lex.Lexer) (*lexer, error) {
	return &lexer{Lexer: lx}, nil
}

func (lx *lexer) position(pos token.Pos) token.Position { return lx.File.Position(pos) }

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
func (lx *lexer) Lex(lval *yySymType) int {
more:
	r := lx.scan()
	if r == '\n' {
		if lx.nest == 0 {
			for _, sym := range yyFollow[lval.yys] {
				if sym == yyEofCode {
					//dbg("%s: EOF in state %v", lx.position(lx.First.Pos()), lval.yys)
					return -1
				}
			}
		}

		//dbg("%s: skip in state %v", lx.position(lx.First.Pos()), lval.yys)
		goto more
	}

	lval.Token = Token{Rune: rune(r), Val: string(lx.TokenBytes(nil))}
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

func (lx *lexer) parse() error {
	if yyParse(lx) != 0 && lx.err == nil {
		return fmt.Errorf("parse: internal error")
	}

	return lx.err
}

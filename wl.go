// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"go/token"
	"io"
	"reflect"

	"github.com/cznic/golex/lex"
	"github.com/cznic/strutil"
)

var (
	testFile *token.File // Testing hook

	hooks = strutil.PrettyPrintHooks{
		reflect.TypeOf(Token{}): func(f strutil.Formatter, v interface{}, prefix string, suffix string) {
			t := v.(Token)
			if t.Rune == 0 {
				return
			}

			f.Format(prefix)
			if testFile != nil {
				f.Format("%s: ", testFile.Position(t.Pos()))
			}
			f.Format("%s", yySymName(int(t.Rune)))
			if t.Val != "" {
				f.Format(", %q", t.Val)
			}
			f.Format(suffix)
		},
	}
)

// Node is implemented by all AST nodes.
type Node interface {
	Pos() token.Pos
}

// Token represents a terminal AST node.
type Token struct {
	Rune rune
	Val  string
	pos  token.Pos
}

// Pos implements Node.
func (t *Token) Pos() token.Pos { return t.pos }

func prettyString(v interface{}) string { return strutil.PrettyString(v, "", "", hooks) }

// Input represents parser's source.
type Input struct {
	interactive bool
	lex         *lex.Lexer
	lookahead   lex.Char
	lx          *lexer
}

// NewInput returns a newly created Input or an error, of any. The interactive
// argument enables ParseExpression to return on newlines in input whenever the
// expression at that point is valid.
func NewInput(r io.RuneReader, interactive bool) (*Input, error) {
	lx := newLexer()
	l, err := lex.New(
		nil,
		r,
		lex.BOMMode(lex.BOMIgnoreFirst),
		lex.RuneClass(runeClass),
		lex.ErrorFunc(lx.errPos),
	)
	if err != nil {
		return nil, err
	}

	p := &Input{lx: lx, lex: l, interactive: interactive}
	p.lookahead.Rune = -1
	return p, nil
}

// ParseExpression parses a Wolfram Language expression using file to record
// and annotate token and error positions, and returns the respective
// *Expression or an error, if any.
func (p *Input) ParseExpression(file *token.File) (*Expression, error) {
	p.lex.File = file
	if p.lookahead.IsValid() && p.lookahead.Rune > 0 {
		p.lex.Unget(p.lookahead)
		p.lookahead = lex.Char{}
	}
	if err := p.lx.parse(p.lex, p.interactive); err != nil {
		return nil, err
	}

	p.lookahead = p.lex.Lookahead()
	return p.lx.ast.(*start).Expression, nil
}

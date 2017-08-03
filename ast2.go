// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"go/token"
	"reflect"

	"github.com/cznic/strutil"
)

var (
	hooks = strutil.PrettyPrintHooks{
		reflect.TypeOf(Token{}): func(f strutil.Formatter, v interface{}, prefix string, suffix string) {
			t := v.(Token)
			if t.Rune == 0 {
				return
			}

			f.Format(prefix)
			f.Format("%s", yySymName(int(t.Rune)))
			if t.Val != "" {
				f.Format(", %q", t.Val)
			}
			f.Format(suffix)
		},
	}
)

type Node interface {
	Pos() token.Pos
}

type Token struct {
	Rune rune
	Val  string
	pos  token.Pos
}

func (t *Token) Pos() token.Pos { return t.pos }

func prettyString(v interface{}) string { return strutil.PrettyString(v, "", "", hooks) }

// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"go/token"

	"github.com/cznic/strutil"
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

func prettyString(v interface{}) string { return strutil.PrettyString(v, "", "", nil) }

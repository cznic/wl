%{
// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
%}

%yyc c
%yyn c = lx.next()

%{
package wl

import (
	"strconv"
)

func (lx *lexer) ReadRune() (r rune, sz int, err error) {
	c := lx.enter()
%}

d	[0-9]
h	{d}|[A-Fa-f]

%%

\\\\	return '\\', 2, lx.rerr
\\\x20	return ' ', 2, lx.rerr
\\\"	return '"', 2, lx.rerr
\\b	return '\b', 2, lx.rerr
\\t	return '\t', 2, lx.rerr
\\n	return '\n', 2, lx.rerr
\\f	return '\f', 2, lx.rerr
\\r	return '\r', 2, lx.rerr

"\["[A-Z][A-Za-z]*"]"
	s := lx.token()
	return lx.named(s[2:len(s)-2])

\\{d}{d}{d}
	n, err := strconv.ParseInt(lx.token()[1:4], 8, 8)
	if err != nil {
		return 0, 0, err
	}

	return rune(n), 4, lx.rerr

\\\.{h}{h}
	n, err := strconv.ParseInt(lx.token()[2:4], 16, 8)
	if err != nil {
		return 0, 0, err
	}

	return rune(n), 4, lx.rerr

\\\:{h}{h}{h}{h}
	n, err := strconv.ParseInt(lx.token()[2:6], 16, 16)
	if err != nil {
		lx.errPos(lx.First.Pos(), err.Error())
		return 0, 0, err
	}

	return rune(n), 6, lx.rerr

%%
	lx.c = -1
	return rune(c), lx.sz, lx.rerr
}

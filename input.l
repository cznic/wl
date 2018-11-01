%{
// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
%}

%yyc c
%yyn c = lx.next()
%yym lx.mark = len(lx.in)

%{
package wl

import (
        "strconv"
        "unicode/utf8"
)

func (lx *lexer) ReadRune() (r rune, sz int, err error) {
        if lx.c < 0 {
                lx.next()
        }
        c := int(lx.c)
%}

d        [0-9]
h        {d}|[A-Fa-f]

%%
        lx.in = lx.in[:0]
        lx.mark = -1

"\["[A-Z][A-Za-z]*"]"
        s := lx.token()
        return lx.named(s[2:len(s)-1])

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
        if lx.mark >= 0 {
                if len(lx.in) > lx.mark {
                        lx.unget(rune(c))
                        for i := len(lx.in)-1; i >= lx.mark; i-- {
                                lx.unget(lx.in[i])
                        }
                        lx.next()
                }
                lx.in = lx.in[:lx.mark]
                goto yyAction
        }
        
        switch n := len(lx.in); n {
        case 0: // [] z
                lx.c = -1
        case 1: // [x] z
                c = int(lx.in[0])
        }

        switch c {
        case '\n':
                lx.unget(IGNORE)
                return rune(c), 1, lx.err
        case IGNORE:
                return rune(c), 0, lx.err
        }

        return rune(c), utf8.RuneLen(rune(c)), lx.err
}

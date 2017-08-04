%{
// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
%}

%yyc c
%yyn c = lx.Next()
%yym lx.Mark()

%{
package wl

import (
	"github.com/cznic/golex/lex"
)

func (lx *lexer) scan() int {
		c := lx.Enter()
%}

/* classes */
EOF		\x80
D		\x81
L		\x82

blank		__?_?
d		[0-9]
decimals	{d}+
exponent	[eE][+-]?{decimals}*
float		{decimals}"."{decimals}?{exponent}?|{decimals}{exponent}|"."{decimals}{exponent}?
ident		{identFirst}{identNext}*
identFirst	[a-zA-Z$]|{L}
identNext	{identFirst}|[0-9]|{D}
int		{decimals}
pattern		{ident}?{blank}{ident}?\.?
slot		#({decimals}|{ident})?

%%
		c = lx.Rule0()
		if lx.err != nil {
			return -1
		}

[ \t\r]+
"(*"([^*\x80]|\*+[^*)\x80])*\*+\)


"&&"			return AND
"->"			return RULE
"/."			return REPLACEALL
"//"			return MAPALL
"//."			return REPLACEREP
"/;"			return CONDITION
"/@"			return MAP
"::"			return MESSAGE
":="			return SET_DELAYED
":>"			return RULEDELAYED
"<="			return LEQ
"<>"			return STRINGJOIN
"=!="			return UNSAME
"=="			return EQUAL
"==="			return SAME
">="			return GEQ
"@@"			return APPLY
"[["			return lx.push(LPART)
"||"			return OR

"]]"			if lx.pop() == LPART {
				return RPART
			}
			if la := lx.Lookahead(); la.Rune != 0 {
				lx.Unget(la)
			}
			lx.Unget(lex.NewChar(lx.First.Pos()+1, ']'))
			return ']'

\"([^"]|\\.)*\"		return STRING
{float}			return FLOAT
{ident}(`{ident})*	return IDENT
{int}			return INT
{pattern}		return PATTERN
{slot}			return SLOT

%%
		if c, ok := lx.Abort(); ok {
			switch c {
			case '(', '[',	'{':
				lx.push(c)
			case ')', ']',	'}':
				lx.pop()
			}
			return c
		}

		goto yyAction
}


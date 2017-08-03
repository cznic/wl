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

func (lx *lexer) scan() int {
		c := lx.Enter()
%}

/* classes */
EOF		\x80
D		\x81
L		\x82

d		[0-9]
decimals	{d}+
exponent	[eE][+-]?{decimals}*
float		{decimals}"."{decimals}?{exponent}?|{decimals}{exponent}|"."{decimals}{exponent}?
ident		{identFirst}{identNext}*
identFirst	[a-zA-Z$]|{L}
identNext	{identFirst}|[0-9]|{D}
int		{decimals}
blank		_|__|___
pattern		{ident}?{blank}{ident}?

%%
		c = lx.Rule0()
		if lx.err != nil {
			return -1
		}

[ \t\r]+

"&&"			return AND
"/;"			return CONDITION
"::"			return MESSAGE
":="			return SET_DELAYED
"=!="			return UNSAME
\"[^"]*\"		return STRING
{float}			return FLOAT
{ident}(`{ident})*	return IDENTIFIER
{int}			return INT
{pattern}		return PATTERN

%%
		if c, ok := lx.Abort(); ok {
			switch c {
			case '[', '{':
				lx.nest++
			case ']', '}':
				lx.nest--
			}
			return c
		}

		goto yyAction
}


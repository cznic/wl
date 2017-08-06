%{
// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
%}

%yyc c
%yyn c = lx.Next()
%yym lx.Mark()
%yyt lx.sc

%x S

%{
package wl

import (
	"fmt"

	"github.com/cznic/golex/lex"
)

func (lx *lexer) scan() (r int) {
	// fmt.Printf("%T.scan\n", lx)
	// defer func() { fmt.Printf("\t%T.scan: %U %s\n", lx, r, yySymName(r)) }()
	c := lx.Enter()
%}

/* classes */
EOF		\x80
D		\x81
L		\x82
LL		\x83
ignore		\x84

approx		{digits}"."{digits}|"."{digits}|{digits}"."
base		{digits}^^
blank		__?_?
d		[0-9]
digiticks	{digits}(`{digits})*
digits	{d}+
exponent	[eE][+-]?{digits}
float		{base}?{approx}(`|`{digits}|``{digits})?
ident		{symbol}|`{symbol}|{symbol}`{symbol}|{symbol}`|{symbol}`{symbol}`|`{symbol}`
int		{base}?{digits}
mantissa	{digiticks}"."{digiticks}|"."{digiticks}|{digiticks}"."
pattern		{symbol}?{blank}{symbol}?\.?
sci		({approx}|{mantissa})"*^"{digits}
slot		#({digits}|{ident})?
symbol		{symbolFirst}{symbolNext}*
symbolFirst	[a-zA-Z$]|{L}|{LL}
symbolNext	{symbolFirst}|[0-9]|{D}
schar		[^\x22]|\n

%%
	c = lx.Rule0()

[ \t\r]+
"(*"([^*\x80]|\*+[^*)\x80])*\*+\)

\"	lx.sc = 1
	lx.str = lx.str[:0]

<S>\x00		lx.str = append(lx.str, 0)
<S>\\\\		lx.str = append(lx.str, '\\')
<S>\\\x20	lx.str = append(lx.str, ' ')
<S>\\\"		lx.str = append(lx.str, '"')
<S>\\b		lx.str = append(lx.str, '\b')
<S>\\t		lx.str = append(lx.str, '\t')
<S>\\n		lx.str = append(lx.str, '\n')
<S>\\f		lx.str = append(lx.str, '\f')
<S>\\r		lx.str = append(lx.str, '\r')
<S>\\\n
<S>{ignore}

<S>{schar}
	lx.str = append(lx.str, lx.TokenBytes(nil)...)

<S>\"	lx.sc = 0
	return STRING

{ignore}
	return IGNORE


"&&"	return AND
"++"	return INC
"--"	return DEC
"->"	return RULE
"/*"	return RIGHT_COMPOSITION
"/."	return REPLACEALL
"//"	return POSTFIX
"//."	return REPLACEREP
"//@"	return MAP_ALL
"/;"	return CONDITION
"/@"	return MAP
"::"	return MESSAGE_NAME
":="	return SET_DELAYED
":>"	return RULEDELAYED
"<<"	return GET
"<="	return LEQ
"<>"	return STRINGJOIN
"=!="	return UNSAME
"=="	return EQUAL
"==="	return SAME
">="	return GEQ
"@*"	return COMPOSITION
"@@"	return APPLY
"@@@"	return APPLY_ALL
"[["	return lx.push(LPART)
"\&"	return OVERSCRIPT
"\+"	return UNDERSCRIPT
"\_"	return SUBSCRIPT
"||"	return OR

"]]"	if lx.pop() == LPART {
		return RPART
	}

	if la := lx.Lookahead(); la.Rune != 0 {
		lx.Unget(la)
	}
	lx.Unget(lex.NewChar(lx.First.Pos()+1, ']'))
	return ']'

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


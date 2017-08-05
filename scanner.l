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

func (lx *lexer) scan() (r int) {
	c := lx.Enter()
%}

/* classes */
EOF		\x80
D		\x81
L		\x82
LL		\x83

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

%%
	c = lx.Rule0()

[ \t\r]+
"(*"([^*\x80]|\*+[^*)\x80])*\*+\)


"&&"	return AND
"->"	return RULE
"/."	return REPLACEALL
"//"	return POSTFIX
"//."	return REPLACEREP
"/;"	return CONDITION
"/@"	return MAP
"::"	return MESSAGE
":="	return SET_DELAYED
":>"	return RULEDELAYED
"<="	return LEQ
"<>"	return STRINGJOIN
"=!="	return UNSAME
"=="	return EQUAL
"==="	return SAME
">="	return GEQ
"@@"	return APPLY
"[["	return lx.push(LPART)
"||"	return OR

"]]"	if lx.pop() == LPART {
		return RPART
	}

	return lx.unget(']')

\"[^"]*\"	return STRING
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


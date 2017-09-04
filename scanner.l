%{
// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
%}

%yyc c
%yyn c = lx.Next()
%yym lx.Mark()
%yyt lx.sc

%x S C C2

%{
package wl

import (
	"fmt"

	"github.com/cznic/golex/lex"
)

func (lx *lexer) scan() (r int) {
	const (
		scInitial = iota
		scString
		scComment
		scComment2
	)
	c := lx.Enter()
%}

/* classes */
EOF			\x80
D			\x81
L			\x82
LL			\x83
ignore			\x84
Conjugate		\x85
Transpose		\x86
ConjugateTranspose	\x87
HermitianConjugate	\x88
Sqrt			\x89
Integrate		\x8a
DifferentialD		\x8b
PartialD		\x8c
Del			\x8d
DiscreteShift		\x8e
DiscreteRatio		\x8f
DifferenceDelta		\x90
Square			\x91
SmallCircle		\x92
CircleDot		\x93
Cross			\x94
PlusMinus		\x95
MinusPlus		\x96
Divide			\x97
Backslash		\x98
Diamond			\x99
Wedge			\x9a
Vee			\x9b
CircleTimes		\x9c
CenterDot		\x9d
Times			\x9e
Star			\x9f
Product			\xa0
VerticalTilde		\xa1
Coproduct		\xa2
Cap			\xa3
Cup			\xa4
CirclePlus		\xa5
CircleMinus		\xa6
Sum			\xa7
Intersection		\xa8
Union			\xa9
Equal			\xaa
VerticalBar		\xab
NotVerticalBar		\xac
DoubleVerticalBar	\xad
NotDoubleVerticalBar	\xae
Element			\xaf
NotElement		\xb0
Subset			\xb1
Superset		\xb2
ForAll			\xb3
Exists			\xb4
NotExists		\xb5
Not			\xb6
And			\xb7
Nand			\xb8
Xor			\xb9
Xnor			\xba
Or			\xbb
Nor			\xbc
Equivalent		\xbd
Implies			\xbe
RightTee		\xbf
DoubleRightTee		\xc0
LeftTee			\xc1
DoubleLeftTee		\xc2
UpTee			\xc3
DownTee			\xc4
SuchThat		\xc5
VerticalSeparator	\xc6
Therefore		\xc7
Because			\xc8
Function		\xc9

approx		{digits}"."{digits}|"."{digits}|{digits}"."
base		{digits}^^
blank		__?_?
d		[0-9]
digiticks	{digits}(`{digits})*
digits	{d}+
exponent	[eE][+-]?{digits}
float		{base}?{approx}(`|`{digits}|``{digits})?
ident		{symbol}|`{symbol}|{symbol}`{symbol}|{symbol}`|{symbol}`{symbol}`|{symbol}`{symbol}`{symbol}|`{symbol}`
int		{base}?{digits}
mantissa	{digiticks}"."{digiticks}|"."{digiticks}|{digiticks}"."
out		%+|%{digits}
pattern		{symbol}?{blank}{symbol}?\.?
schar		[^\x22]|\n
sci		({approx}|{mantissa})"*^"{digits}
slot		#+|#({digits}|{ident})?
symbol		{symbolFirst}{symbolNext}*
symbolFirst	[a-zA-Z$]|{L}|{LL}
symbolNext	{symbolFirst}|[0-9]|{D}

%%
	c = lx.Rule0()

[ \t\r]+

"(*"		lx.sc = scComment
		lx.commentLevel++

<C>"(*"		lx.commentLevel++

<C>"*"		lx.sc = scComment2

<C>.|\n

<C2>"(*"	lx.commentLevel++

<C2>"*"

<C2>")"		lx.commentLevel--
		if lx.commentLevel == 0 {
			lx.sc = scInitial
		}

<C2>.|\n
		lx.sc = scComment


\"		lx.sc = scString
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

<S>\"	lx.sc = scInitial
	return STRING

{ignore}
	return IGNORE

"!="	return UNEQUAL
"&&"	return AND
"'"	return QUOTE
"**"	return NON_COMMUTATIVE_MULTIPLY
"*="	return TIMES_BY
"++"	return INC
"+="	return ADD_TO
"--"	return DEC
"-="	return SUBTRACT_FROM
"->"	return RULE
".."	return REPEATED
"..."	return REPEATED_NULL
"/*"	return RIGHT_COMPOSITION
"/."	return REPLACEALL
"//"	return POSTFIX
"//."	return REPLACEREP
"//@"	return MAP_ALL
"/:"	return TAG_SET
"/;"	return CONDITION
"/="	return DIVIDE_BY
"/@"	return MAP
"::"	return MESSAGE_NAME
":="	return SET_DELAYED
":>"	return RULEDELAYED
";;"	return SPAN
"<<"	return GET
"<="	return LEQ
"<>"	return STRINGJOIN
"=!="	return UNSAME
"=="	return EQUAL
"==="	return SAME
">="	return GEQ
">>"	return PUT
">>>"	return PUT_APPEND
"@*"	return COMPOSITION
"@@"	return APPLY
"@@@"	return APPLY_ALL
"[["	return lx.push(LPART)
"\%"	return POWER_SUBSCRIPT2
"\&"	return OVERSCRIPT
"\+"	return UNDERSCRIPT
"\/"	return '/'
"\@"	return SQRT2
"\^"	return POWER_SUBSCRIPT1
"\_"	return SUBSCRIPT
"\`"	return FORM_BOX
"^:="	return UP_SET_DELAYED
"^="	return UP_SET
"||"	return OR
"~~"	return STRING_EXPRESSION

"]]"	if lx.pop() == LPART {
		return RPART
	}

	if la := lx.Lookahead(); la.Rune != 0 {
		lx.Unget(la)
	}
	lx.Unget(lex.NewChar(lx.First.Pos()+1, ']'))
	return ']'

{And}			return AND
{Backslash}		return BACKSLASH
{Because}		return BECAUSE
{Cap}			return CAP
{CenterDot}		return CENTER_DOT
{CircleDot}		return CIRCLE_DOT
{CircleMinus}		return CIRCLE_MINUS
{CirclePlus}		return CIRCLE_PLUS
{CircleTimes}		return CIRCLE_TIMES
{ConjugateTranspose}	return CONJUGATE_TRANSPOSE
{Conjugate}		return CONJUGATE
{Coproduct}		return COPRODUCT
{Cross}			return CROSS
{Cup}			return CUP
{Del}			return DEL
{Diamond}		return DIAMOND
{DifferenceDelta}	return DIFFERENCE_DELTA
{DifferentialD}		return DIFFERENTIAL_D
{DiscreteRatio}		return DISCRETE_RATIO
{DiscreteShift}		return DISCRETE_SHIFT
{Divide}		return '/'
{DoubleLeftTee}		return DOUBLE_LEFT_TEE
{DoubleRightTee}	return DOUBLE_RIGHT_TEE
{DoubleVerticalBar}	return DOUBLE_VERTICAL_BAR
{DownTee}		return DOWN_TEE
{Element}		return ELEMENT
{Equal}			return EQUAL
{Equivalent}		return EQUIVALENT
{Exists}		return EXISTS
{ForAll}		return FOR_ALL
{Function}		return FUNCTION
{HermitianConjugate}	return HERMITIAN_CONJUGATE
{Implies}		return IMPLIES
{Integrate}		return INTEGRATE
{Intersection}		return INTERSECTION
{LeftTee}		return LEFT_TEE
{MinusPlus}		return MINUS_PLUS
{Nand}			return NAND
{Nor}			return NOR
{NotDoubleVerticalBar}	return NOT_DOUBLE_VERTICAL_BAR
{NotElement}		return NOT_ELEMENT
{NotExists}		return NOT_EXISTS
{NotVerticalBar}	return NOT_VERTICAL_BAR
{Not}			return '!'
{Or}			return OR
{PartialD}		return PARTIAL_D
{PlusMinus}		return PLUS_MINUS
{Product}		return PRODUCT
{RightTee}		return RIGHT_TEE
{SmallCircle}		return SMALL_CIRCLE
{Sqrt}			return SQRT
{Square}		return SQUARE
{Star}			return STAR
{Subset}		return SUBSET
{SuchThat}		return SUCH_THAT
{Sum}			return SUM
{Superset}		return SUPERSET
{Therefore}		return THEREFORE
{Times}			return '*'
{Transpose}		return TRANSPOSE
{Union}			return UNION
{UpTee}			return UP_TEE
{Vee}			return VEE
{VerticalBar}		return VERTICAL_BAR
{VerticalSeparator}	return VERTICAL_SEPARATOR
{VerticalTilde}		return VERTICAL_TILDE
{Wedge}			return WEDGE
{Xnor}			return XNOR
{Xor}			return XOR

{float}		return FLOAT
{ident}		return IDENT
{int}		return INT
{out}		return OUT
{pattern}	return PATTERN
{slot}		return SLOT

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


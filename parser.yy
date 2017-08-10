%{
// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

%}

%union{
	Node Node
	Token Token
}

%token
	/*yy:token "#%d"    */	SLOT		"slot"
	/*yy:token "%c"     */	IDENT		"identifier"
	/*yy:token "%c_"    */	PATTERN		"pattern"
	/*yy:token "%d"     */	INT		"integer"
	/*yy:token "1.%d"   */	FLOAT		"real"
	/*yy:token "\"%c\"" */	STRING		"string"

%token
	AND				"&&"
	APPLY				"@@"
	APPLY_ALL			"@@@"
	BACKSLASH			"\\[Backslash]"
	CENTER_DOT			"\\[CenterDot]"
	CIRCLE_DOT			"\\[CircleDot]"
	CIRCLE_TIMES			"\\[CircleTimes]"
	COMPOSITION			"@*"
	CONDITION			"/;"
	CONJUGATE			"\\[Conjugate]"
	CONJUGATE_TRANSPOSE		"\\[ConjugateTranspose]"
	CROSS				"\\[Cross]"
	DEC				"--"
	DEL				"\\[Del]"
	DIAMOND				"\\[Diamond]"
	DIFFERENCE_DELTA		"\\[DifferenceDelta]"
	DIFFERENTIAL_D			"\\[DifferentialD]"
	DISCRETE_RATIO			"\\[DiscreteRatio]"
	DISCRETE_SHIFT			"\\[DiscreteShift]"
	DIVIDE				"\\[Divide]"
	DIVIDE2				"\\/"
	EQUAL				"=="
	GEQ				">="
	GET				"<<"
	HERMITIAN_CONJUGATE		"\\[HermitianConjugate]"
	INC				"++"
	INTEGRATE			"\\[Integrate]"
	LEQ				"<="
	LPART				"[["
	MAP				"/@"
	MAP_ALL				"//@"
	MESSAGE_NAME			"::"
	MINUS_PLUS			"\\[MinusPlus]"
	NON_COMMUTATIVE_MULTIPLY	"**"
	OR				"||"
	OVERSCRIPT			"\\&"
	PARTIAL_D			"\\[PartialD]"
	PLUS_MINUS			"\\[PlusMinus]"
	POSTFIX				"//"
	POWER_SUBSCRIPT1		"\\^"
	POWER_SUBSCRIPT2		"\\%"
	QUOTE				"'"
	REPLACEALL			"/."
	REPLACEREP			"//."
	RIGHT_COMPOSITION		"/*"
	RPART				"]]"
	RULE				"->"
	RULEDELAYED			":>"
	SAME				"==="
	SET_DELAYED			":="
	SMALL_CIRCLE			"\\[SmallCircle]"
	SQRT				"\\[Sqrt]"
	SQRT2				"\\@"
	SQUARE				"\\[Square]"
	STRINGJOIN			"<>"
	SUBSCRIPT			"\\_"
	TRANSPOSE			"\\[Transpose]"
	UNDERSCRIPT			"\\+"
	UNEQUAL				"!="
	UNSAME				"=!="
	VEE				"\\[Vee]"
	WEDGE				"\\[Wedge]"

%type	<Node>
	start		"valid input"
	CommaOpt	"optional comma"
	ExprList	"expression list"
	Expression	"expression"
	Factor		"factor"
	Tag		"tag"
	Term		"term"

%token IGNORE

%left ';'
%left '=' SET_DELAYED
%left POSTFIX
%left ':'
%precedence '&'
%left REPLACEREP
%left REPLACEALL
%left RULEDELAYED
%left RULE
%left CONDITION
%left PATTERN
%left '|'
%left OR
%left AND
%right '!'			// Not
%left UNSAME
%left SAME
%left LEQ
%left '<'
%left GEQ
%left '>'
%left UNEQUAL
%left EQUAL
%left '+'
%left '-'
%left '*'

%left		CENTER_DOT
%left		CIRCLE_TIMES
%left		VEE
%left		WEDGE
%left		DIAMOND
%left		BACKSLASH
%left		'/'
%precedence	UNARY_MINUS UNARY_PLUS PLUS_MINUS MINUS_PLUS
%left		'.'			// Dot
%right		CROSS
%left		NON_COMMUTATIVE_MULTIPLY
%right		CIRCLE_DOT
%right		SQUARE SMALL_CIRCLE
%right		PARTIAL_D DEL DISCRETE_SHIFT DISCRETE_RATIO DIFFERENCE_DELTA
%right		INTEGRATE DIFFERENTIAL_D
%right		SQRT SQRT2
%right 		'^' POWER_SUBSCRIPT1 POWER_SUBSCRIPT2	// Power, Power[Subscript]
%left		STRINGJOIN
%nonassoc	QUOTE
%nonassoc	CONJUGATE TRANSPOSE CONJUGATE_TRANSPOSE HERMITIAN_CONJUGATE
%nonassoc	FACTORIAL
%right		MAP MAP_ALL APPLY APPLY_ALL
%left		'~'
%right		'@'
%left		COMPOSITION RIGHT_COMPOSITION
%precedence	PRE_INC		// PreIncrement, PreDecrement
%nonassoc	INC DEC		// Increment, Decrement
%left		'[' ']' LPART RPART	// expr, Part
%left		'?'	// PatternTest
%right		SUBSCRIPT
%right		OVERSCRIPT UNDERSCRIPT
%nonassoc	GET
/*TODO forms containing # */
%left		MESSAGE_NAME
/*TODO Piecewise */

%%

start:
	Expression

Expression:
	"++" Expression %prec PRE_INC
|	"--" Expression %prec PRE_INC
|	"\\@" Expression
|	"\\@" Expression "\\%" Expression
|	"\\[Sqrt]" Expression
|	"\\[Integrate]" Expression "\\[DifferentialD]" Expression
|	'!' Expression
|	'-' Expression %prec UNARY_MINUS
|	Expression "&&" Expression
|	Expression "++"
|	Expression "--"
|	Expression "->" Expression
|	Expression "/*" Expression
|	Expression "/." Expression
|	Expression "//" Expression
|	Expression "//." Expression
|	Expression "//@" Expression
|	Expression "/;" Expression
|	Expression "/@" Expression
|	Expression ":=" Expression
|	Expression ":>" Expression
|	Expression "<=" Expression
|	Expression "<>" Expression
|	Expression "=!=" Expression
|	Expression "==" Expression
|	Expression "===" Expression
|	Expression ">=" Expression
|	Expression "@*" Expression
|	Expression "@@" Expression
|	Expression "@@@" Expression
|	Expression "\\&" Expression
|	Expression "\\+" Expression
|	Expression "\\^" Expression "\\%" Expression
|	Expression "\\_" Expression
|	Expression "||" Expression
|	Expression '*' Expression
|	Expression '+' Expression
|	Expression '-' Expression
|	Expression '.' Expression
|	Expression '/' Expression
|	Expression ':' Expression
|	Expression ';'
|	Expression ';' Expression
|	Expression '<' Expression
|	Expression '=' Expression
|	Expression '>' Expression
|	Expression '?' Expression
|	Expression '@' Expression
|	Expression '^' Expression
|	Expression '|' Expression
|	Expression '~' Expression
|	Expression "\\[Conjugate]"
|	Expression "\\[ConjugateTranspose]"
|	Expression "\\[HermitianConjugate]"
|	Expression "\\[Transpose]"
|	Factor
|	Expression "\\[PartialD]" Expression
|	"\\[Del]" Expression
|	Expression "\\[DiscreteShift]" Expression
|	Expression "\\[DiscreteRatio]" Expression
|	Expression "\\[DifferenceDelta]" Expression
|	"\\[Square]" Expression
|	Expression "\\[SmallCircle]" Expression
|	Expression "\\[CircleDot]" Expression
|	Expression "**" Expression
|	Expression "\\[Cross]" Expression
|	'+' Expression %prec UNARY_PLUS
|	"\\[PlusMinus]" Expression
|	"\\[MinusPlus]" Expression
|	Expression "\\[Backslash]" Expression
|	Expression "!=" Expression
|	Expression "\\[Diamond]" Expression
|	Expression "\\[Wedge]" Expression
|	Expression "\\[Vee]" Expression
|	Expression "\\[CircleTimes]" Expression
|	Expression "\\[CenterDot]" Expression

Term:
	"<<" STRING
|	'(' Expression ')'
|	'{' '}'
|	'{' ExprList CommaOpt '}'
|	FLOAT
|	IDENT
|	IDENT "::" Tag
|	IDENT "::" Tag "::" Tag
|	INT
|	PATTERN
|	SLOT
|	STRING
|	Term "[[" ExprList CommaOpt "]]"
|	Term '!'
|	Term '!' '!'
|	Term '&'
|	Term '[' ']'
|	Term '[' ExprList CommaOpt ']'
|	Term QUOTE

Factor:
	Term
|	Term Factor

ExprList:
	Expression
|	ExprList ',' Expression

CommaOpt:
	/* empty */ {}
|	','

Tag:
	IDENT
|	STRING

%%

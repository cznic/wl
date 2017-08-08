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
	COMPOSITION			"@*"
	CONDITION			"/;"
	DEC				"--"
	DIVIDE2				"\\/"
	EQUAL				"=="
	UNEQUAL				"!="
	GEQ				">="
	GET				"<<"
	INC				"++"
	INTEGRATE			"∫"
	LEQ				"<="
	LPART				"[["
	MAP				"/@"
	MAP_ALL				"//@"
	MESSAGE_NAME			"::"
	NON_COMMUTATIVE_MULTIPLY	"**"
	OR				"||"
	OVERSCRIPT			"\\&"
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
	SQRT				"√"
	SQRT2				"\\@"
	STRINGJOIN			"<>"
	SUBSCRIPT			"\\_"
	UNDERSCRIPT			"\\+"
	UNSAME				"=!="

	BACKSLASH			"\u2216"
	CIRCLE_DOT			"\u2299"
	CONJUGATE			"\uf3c8"
	CONJUGATE_TRANSPOSE		"\uf3c9"
	CROSS				"\uf4a0"
	DEL				"\u2207"
	DIFFERENCE_DELTA		"\u2206"
	DIFFERENTIAL_D			"\uf74c"
	DISCRETE_RATIO			"\uf4a4"
	DISCRETE_SHIFT			"\uf4a3"
	DIVIDE				"\u00f7"
	HERMITIAN_CONJUGATE		"\uf3ce"
	MINUS_PLUS			"\u2213"
	PARTIAL_D			"\u2202"
	PLUS_MINUS			"\u00b1"
	SMALL_CIRCLE			"\u2218"
	SQUARE				"\uf520"
	TRANSPOSE			"\uf3c7"

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

%left BACKSLASH
%left '/'
%precedence UNARY_MINUS UNARY_PLUS PLUS_MINUS MINUS_PLUS
%left	'.'			// Dot
%right	CROSS
%left	NON_COMMUTATIVE_MULTIPLY
%right	CIRCLE_DOT
%right	SQUARE SMALL_CIRCLE
%right	PARTIAL_D DEL DISCRETE_SHIFT DISCRETE_RATIO DIFFERENCE_DELTA
%right	INTEGRATE DIFFERENTIAL_D
%right	SQRT SQRT2
%right 	'^' POWER_SUBSCRIPT1 POWER_SUBSCRIPT2	// Power, Power[Subscript]
%left STRINGJOIN
%nonassoc	QUOTE
%nonassoc	CONJUGATE TRANSPOSE CONJUGATE_TRANSPOSE HERMITIAN_CONJUGATE
%nonassoc	FACTORIAL
%right	MAP MAP_ALL APPLY APPLY_ALL
%left	'~'
%right	'@'
%left	COMPOSITION RIGHT_COMPOSITION
%precedence	PRE_INC		// PreIncrement, PreDecrement
%nonassoc	INC DEC		// Increment, Decrement
%left	'[' ']' LPART RPART	// expr, Part
%left	'?'	// PatternTest
%right	SUBSCRIPT
%right	OVERSCRIPT UNDERSCRIPT
%nonassoc	GET
/*TODO forms containing # */
%left	MESSAGE_NAME
/* TODO Piecewise */

%%

start:
	Expression

Expression:
	"++" Expression %prec PRE_INC
|	"--" Expression %prec PRE_INC
|	"\\@" Expression
|	"\\@" Expression "\\%" Expression
|	"√" Expression
|	"∫" Expression DIFFERENTIAL_D Expression
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
|	Expression CONJUGATE
|	Expression CONJUGATE_TRANSPOSE
|	Expression HERMITIAN_CONJUGATE
|	Expression TRANSPOSE
|	Factor
|	Expression PARTIAL_D Expression
|	DEL Expression
|	Expression DISCRETE_SHIFT Expression
|	Expression DISCRETE_RATIO Expression
|	Expression DIFFERENCE_DELTA Expression
|	SQUARE Expression
|	Expression SMALL_CIRCLE Expression
|	Expression CIRCLE_DOT Expression
|	Expression NON_COMMUTATIVE_MULTIPLY Expression
|	Expression CROSS Expression
|	'+' Expression %prec UNARY_PLUS
|	PLUS_MINUS Expression
|	MINUS_PLUS Expression
|	Expression BACKSLASH Expression
|	Expression "!=" Expression

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

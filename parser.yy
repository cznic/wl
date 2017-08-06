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
	AND			"&&"
	APPLY			"@@"
	APPLY_ALL		"@@@"
	COMPOSITION		"@*"
	CONDITION		"/;"
	DEC			"--"
	EQUAL			"=="
	GEQ			">="
	GET			"<<"
	INC			"++"
	LEQ			"<="
	LPART			"[["
	MAP			"/@"
	MAP_ALL			"//@"
	MESSAGE_NAME		"::"
	OR			"||"
	OVERSCRIPT		"\\&"
	POSTFIX			"//"
	POWER_SUBSCRIPT1	"\\^"
	POWER_SUBSCRIPT2	"\\%"
	QUOTE			"'"
	REPLACEALL		"/."
	REPLACEREP		"//."
	RIGHT_COMPOSITION	"/*"
	RPART			"]]"
	RULE			"->"
	RULEDELAYED		":>"
	SAME			"==="
	SET_DELAYED		":="
	SQRT			"√"
	SQRT2			"\\@"
	STRINGJOIN		"<>"
	SUBSCRIPT		"\\_"
	UNDERSCRIPT		"\\+"
	UNSAME			"=!="
	INTEGRATE		"∫"

	CONJUGATE			"\uf3c8"
	CONJUGATE_TRANSPOSE		"\uf3c9"
	DIFFERENTIAL_D			"\uf74c"
	HERMITIAN_CONJUGATE		"\uf3ce"
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
%left EQUAL
%left '-'
%left '+'
%left '*'
%left '/'
%precedence UNARYMINUS
%left '.'

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
|	'-' Expression %prec UNARYMINUS
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

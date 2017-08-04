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
	/*yy:token "%d"     */	INT		"integer literal"
	/*yy:token "1.%d"   */	FLOAT		"floating point literal"
	/*yy:token "\"%c\"" */	STRING		"string literal"

%token
	AND		"&&"
	APPLY		"@@"
	CONDITION	"/;"
	EQUAL		"=="
	GEQ		">="
	LEQ		"<="
	LPART		"[["
	MAP		"/@"
	MAPALL		"//"
	MESSAGE		"::"
	OR		"||"
	REPLACEALL	"/."
	REPLACEREP	"//."
	RPART		"]]"
	RULE		"->"
	RULEDELAYED	":>"
	SAME		"==="
	SET_DELAYED	":="
	STRINGJOIN	"<>"
	UNSAME		"=!="

%left ';'
%left '=' SET_DELAYED
%left ':'
%precedence '&'
%left REPLACEREP
%left REPLACEALL
%left RULEDELAYED
%left RULE
%left CONDITION
%precedence NOPATTERN
%left PATTERN
%left '|'
%left OR
%left AND
%precedence NOT
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
%right '^'
%left STRINGJOIN
%precedence '!' // Factorial
%left APPLY
%left MAPALL
%left MAP
%left '@'
%left LPART RPART
%left '[' ']'
%left '?'

%%

start:
	Expression

Expression:
	'!' Expression %prec NOT
|	'-' Expression %prec UNARYMINUS
|	Expression "&&" Expression
|	Expression "->" Expression
|	Expression "/." Expression
|	Expression "//" Expression
|	Expression "//." Expression
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
|	Expression "@@" Expression
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
|	Factor %prec NOPATTERN
|	Factor ':' Expression %prec PATTERN

Term:
	FLOAT
|	'(' Expression ')'
|	'{' '}'
|	'{' ExprList CommaOpt '}'
|	IDENT
|	IDENT "::" Tag
|	IDENT "::" Tag "::" Tag
|	INT
|	PATTERN
|	SLOT
|	STRING
|	Term "[[" ExprList CommaOpt "]]"
|	Term '!'
|	Term '&'
|	Term '[' ']'
|	Term '[' ExprList CommaOpt ']'

Factor:
	Term
|	Term Factor

ExprList:
	Expression
|	ExprList ',' Expression

CommaOpt:
|	','

Tag:
	IDENT
|	STRING

%%

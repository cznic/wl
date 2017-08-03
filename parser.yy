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
	/*yy:token "%c"     */	IDENTIFIER	"identifier"
	/*yy:token "%d"     */	INT		"integer literal"
	/*yy:token "1.%d"   */	FLOAT		"floating point literal"
	/*yy:token "\"%c\"" */	STRING		"string literal"
	/*yy:token "_"      */	PATTERN		"pattern"

%token
	AND		"&&"
	APPLY		"@@"
	CONDITION	"/;"
	EQUAL		"=="
	GEQ		">="
	LEQ		"<="
	MAP		"/@"
	MAPALL		"//"
	MESSAGE		"::"
	REPLACEALL	"/."
	RULE		"->"
	SAME		"==="
	SET_DELAYED	":="
	SLOT		"#"
	UNSAME		"=!="

%left '=' SET_DELAYED
%precedence '&'
%left REPLACEALL
%left RULE
%left CONDITION
%left AND
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
%precedence UNARY
%right '^'
%precedence '!'
%left APPLY
%left MAPALL
%left MAP
%precedence '[' ']'
%left '?'

%%

Start:
	Expr
|	Expr ';'

Term:
	"#"
|	"identifier"
|	"integer literal"
|	"pattern"
|	"string literal"
|	'(' Expr ')'
|	Term '&'
|	Term '[' ']'
|	Term '[' ExprList ']'
| 	"floating point literal"

Factor:
	Term
|	Term Factor

Expr:
	"identifier" "::" Tag "::" Tag '=' Expr
|	"identifier" "::" Tag '=' Expr
|	'-' Expr %prec UNARY
|	'{' '}'
|	'{' ExprList '}'
|	Expr "&&" Expr
|	Expr "->" Expr
|	Expr "/." Expr
|	Expr "//" Expr
|	Expr "/;" Expr
|	Expr "/@" Expr
|	Expr ":=" Expr
|	Expr "<=" Expr
|	Expr "=!=" Expr
|	Expr "==" Expr
|	Expr "===" Expr
|	Expr ">=" Expr
|	Expr "@@" Expr
|	Expr '!'
|	Expr '*' Expr
|	Expr '+' Expr
|	Expr '-' Expr
|	Expr '/' Expr
|	Expr '<' Expr
|	Expr '=' Expr
|	Expr '>' Expr
|	Expr '?' Expr
|	Expr '^' Expr
|	Factor

ExprList:
	Expr
|	ExprList ',' Expr

Tag:
	"identifier"
|	"string literal"

%%

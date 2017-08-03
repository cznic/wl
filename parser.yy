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
	CONDITION	"/;"
	MESSAGE		"::"
	SET_DELAYED	":="
	UNSAME		"=!="

%left CONDITION
%left UNSAME
%left AND
%left '+'
%left '*' '/'
%left '!'
%right	'=' SET_DELAYED
%precedence UNARY

%%

Start:
	Expr
|	Expr ';'

Expr:
 	"floating point literal"
|	"identifier"
|	"identifier" "::" Tag "::" Tag '=' Expr
|	"identifier" "::" Tag '=' Expr
|	"identifier" '[' ExprList ']'
|	"integer literal"
|	"pattern"
|	"string literal"
|	'(' Expr ')'
|	'-' Expr %prec UNARY
|	'{' '}'
|	'{' ExprList '}'
|	Expr "=!=" Expr
|	Expr "&&" Expr
|	Expr "/;" Expr
|	Expr ":=" Expr
|	Expr '!'
|	Expr '*' Expr
|	Expr '+' Expr
|	Expr '/' Expr
|	Expr '=' Expr

ExprList:
	Expr
|	ExprList ',' Expr

Tag:
	"identifier"
|	"string literal"

%%

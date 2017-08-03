// Code generated by yy. DO NOT EDIT.

// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"go/token"
)

// Expr represents data reduced by productions:
//
//	Expr:
//	        "identifier" "::" Tag "::" Tag '=' Expr
//	|       "identifier" "::" Tag '=' Expr           // Case 1
//	|       '-' Expr                                 // Case 2
//	|       '{' '}'                                  // Case 3
//	|       '{' ExprList '}'                         // Case 4
//	|       Expr "&&" Expr                           // Case 5
//	|       Expr "->" Expr                           // Case 6
//	|       Expr "/." Expr                           // Case 7
//	|       Expr "//" Expr                           // Case 8
//	|       Expr "/;" Expr                           // Case 9
//	|       Expr "/@" Expr                           // Case 10
//	|       Expr ":=" Expr                           // Case 11
//	|       Expr "<=" Expr                           // Case 12
//	|       Expr "=!=" Expr                          // Case 13
//	|       Expr "==" Expr                           // Case 14
//	|       Expr "===" Expr                          // Case 15
//	|       Expr ">=" Expr                           // Case 16
//	|       Expr "@@" Expr                           // Case 17
//	|       Expr '!'                                 // Case 18
//	|       Expr '*' Expr                            // Case 19
//	|       Expr '+' Expr                            // Case 20
//	|       Expr '-' Expr                            // Case 21
//	|       Expr '/' Expr                            // Case 22
//	|       Expr '<' Expr                            // Case 23
//	|       Expr '=' Expr                            // Case 24
//	|       Expr '>' Expr                            // Case 25
//	|       Expr '?' Expr                            // Case 26
//	|       Expr '^' Expr                            // Case 27
//	|       Factor                                   // Case 28
type Expr struct {
	Case     int
	Expr     *Expr
	Expr2    *Expr
	ExprList *ExprList
	Factor   *Factor
	Tag      *Tag
	Tag2     *Tag
	Token    Token
	Token2   Token
	Token3   Token
	Token4   Token
}

func (n *Expr) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Expr) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Expr) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27:
		return n.Expr.Pos()
	case 28:
		return n.Factor.Pos()
	case 0, 1, 2, 3, 4:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// ExprList represents data reduced by productions:
//
//	ExprList:
//	        Expr
//	|       ExprList ',' Expr  // Case 1
type ExprList struct {
	Case     int
	Expr     *Expr
	ExprList *ExprList
	Token    Token
}

func (n *ExprList) reverse() *ExprList {
	if n == nil {
		return nil
	}

	na := n
	nb := na.ExprList
	for nb != nil {
		nc := nb.ExprList
		nb.ExprList = na
		na = nb
		nb = nc
	}
	n.ExprList = nil
	return na
}

func (n *ExprList) fragment() interface{} { return n.reverse() }

// String implements fmt.Stringer.
func (n *ExprList) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *ExprList) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 0:
		return n.Expr.Pos()
	case 1:
		return n.ExprList.Pos()
	default:
		panic("internal error")
	}
}

// Factor represents data reduced by productions:
//
//	Factor:
//	        Term
//	|       Term Factor  // Case 1
type Factor struct {
	Case   int
	Factor *Factor
	Term   *Term
}

func (n *Factor) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Factor) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Factor) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Term.Pos()
}

// Start represents data reduced by productions:
//
//	Start:
//	        Expr
//	|       Expr ';'  // Case 1
type Start struct {
	Case  int
	Expr  *Expr
	Token Token
}

func (n *Start) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Start) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Start) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Expr.Pos()
}

// Tag represents data reduced by productions:
//
//	Tag:
//	        "identifier"
//	|       "string literal"  // Case 1
type Tag struct {
	Case  int
	Token Token
}

func (n *Tag) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Tag) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Tag) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// Term represents data reduced by productions:
//
//	Term:
//	        "#"
//	|       "identifier"              // Case 1
//	|       "integer literal"         // Case 2
//	|       "pattern"                 // Case 3
//	|       "string literal"          // Case 4
//	|       '(' Expr ')'              // Case 5
//	|       Term '&'                  // Case 6
//	|       Term '[' ']'              // Case 7
//	|       Term '[' ExprList ']'     // Case 8
//	|       "floating point literal"  // Case 9
type Term struct {
	Case     int
	Expr     *Expr
	ExprList *ExprList
	Term     *Term
	Token    Token
	Token2   Token
}

func (n *Term) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Term) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Term) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 6, 7, 8:
		return n.Term.Pos()
	case 0, 1, 2, 3, 4, 5, 9:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}
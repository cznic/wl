// Code generated by yy. DO NOT EDIT.

// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"go/token"
)

// CommaOpt represents data reduced by productions:
//
//	CommaOpt:
//	        /* empty */
//	|       ','          // Case 1
type CommaOpt struct {
	Token Token
}

func (n *CommaOpt) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *CommaOpt) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *CommaOpt) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// ExprList represents data reduced by productions:
//
//	ExprList:
//	        Expression
//	|       ExprList ',' Expression  // Case 1
type ExprList struct {
	Case       int
	ExprList   *ExprList
	Expression *Expression
	Token      Token
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
	case 1:
		return n.ExprList.Pos()
	case 0:
		return n.Expression.Pos()
	default:
		panic("internal error")
	}
}

// Expression represents data reduced by productions:
//
//	Expression:
//	        "++" Expression
//	|       "--" Expression                                            // Case 1
//	|       "\\@" Expression                                           // Case 2
//	|       "\\@" Expression "\\%" Expression                          // Case 3
//	|       "\\[Sqrt]" Expression                                      // Case 4
//	|       "\\[Integrate]" Expression "\\[DifferentialD]" Expression  // Case 5
//	|       '!' Expression                                             // Case 6
//	|       '-' Expression                                             // Case 7
//	|       Expression "&&" Expression                                 // Case 8
//	|       Expression "++"                                            // Case 9
//	|       Expression "--"                                            // Case 10
//	|       Expression "->" Expression                                 // Case 11
//	|       Expression "/*" Expression                                 // Case 12
//	|       Expression "/." Expression                                 // Case 13
//	|       Expression "//" Expression                                 // Case 14
//	|       Expression "//." Expression                                // Case 15
//	|       Expression "//@" Expression                                // Case 16
//	|       Expression "/;" Expression                                 // Case 17
//	|       Expression "/@" Expression                                 // Case 18
//	|       Expression ":=" Expression                                 // Case 19
//	|       Expression ":>" Expression                                 // Case 20
//	|       Expression "<=" Expression                                 // Case 21
//	|       Expression "<>" Expression                                 // Case 22
//	|       Expression "=!=" Expression                                // Case 23
//	|       Expression "==" Expression                                 // Case 24
//	|       Expression "===" Expression                                // Case 25
//	|       Expression ">=" Expression                                 // Case 26
//	|       Expression "@*" Expression                                 // Case 27
//	|       Expression "@@" Expression                                 // Case 28
//	|       Expression "@@@" Expression                                // Case 29
//	|       Expression "\\&" Expression                                // Case 30
//	|       Expression "\\+" Expression                                // Case 31
//	|       Expression "\\^" Expression "\\%" Expression               // Case 32
//	|       Expression "\\_" Expression                                // Case 33
//	|       Expression "||" Expression                                 // Case 34
//	|       Expression '*' Expression                                  // Case 35
//	|       Expression '+' Expression                                  // Case 36
//	|       Expression '-' Expression                                  // Case 37
//	|       Expression '.' Expression                                  // Case 38
//	|       Expression '/' Expression                                  // Case 39
//	|       PATTERN ':' Term                                           // Case 40
//	|       Expression ';'                                             // Case 41
//	|       Expression ';' Expression                                  // Case 42
//	|       Expression '<' Expression                                  // Case 43
//	|       Expression '=' Expression                                  // Case 44
//	|       Expression '>' Expression                                  // Case 45
//	|       Expression '?' Expression                                  // Case 46
//	|       Expression '@' Expression                                  // Case 47
//	|       Expression '^' Expression                                  // Case 48
//	|       Expression '|' Expression                                  // Case 49
//	|       Expression '~' Expression                                  // Case 50
//	|       Expression "\\[Conjugate]"                                 // Case 51
//	|       Expression "\\[ConjugateTranspose]"                        // Case 52
//	|       Expression "\\[HermitianConjugate]"                        // Case 53
//	|       Expression "\\[Transpose]"                                 // Case 54
//	|       Factor                                                     // Case 55
//	|       Expression "\\[PartialD]" Expression                       // Case 56
//	|       "\\[Del]" Expression                                       // Case 57
//	|       Expression "\\[DiscreteShift]" Expression                  // Case 58
//	|       Expression "\\[DiscreteRatio]" Expression                  // Case 59
//	|       Expression "\\[DifferenceDelta]" Expression                // Case 60
//	|       "\\[Square]" Expression                                    // Case 61
//	|       Expression "\\[SmallCircle]" Expression                    // Case 62
//	|       Expression "\\[CircleDot]" Expression                      // Case 63
//	|       Expression "**" Expression                                 // Case 64
//	|       Expression "\\[Cross]" Expression                          // Case 65
//	|       '+' Expression                                             // Case 66
//	|       "\\[PlusMinus]" Expression                                 // Case 67
//	|       "\\[MinusPlus]" Expression                                 // Case 68
//	|       Expression "\\[Backslash]" Expression                      // Case 69
//	|       Expression "!=" Expression                                 // Case 70
//	|       Expression "\\[Diamond]" Expression                        // Case 71
//	|       Expression "\\[Wedge]" Expression                          // Case 72
//	|       Expression "\\[Vee]" Expression                            // Case 73
//	|       Expression "\\[CircleTimes]" Expression                    // Case 74
//	|       Expression "\\[CenterDot]" Expression                      // Case 75
//	|       Expression "\\[Star]" Expression                           // Case 76
//	|       Expression "\\[VerticalTilde]" Expression                  // Case 77
//	|       Expression "\\[Coproduct]" Expression                      // Case 78
//	|       Expression "\\[Cap]" Expression                            // Case 79
//	|       Expression "\\[Cup]" Expression                            // Case 80
//	|       Expression "\\[CirclePlus]" Expression                     // Case 81
//	|       Expression "\\[CircleMinus]" Expression                    // Case 82
//	|       Expression "\\[Intersection]" Expression                   // Case 83
//	|       Expression "\\[Union]" Expression                          // Case 84
//	|       ";;"                                                       // Case 85
//	|       ";;" Expression                                            // Case 86
//	|       Expression ";;"                                            // Case 87
//	|       Expression ";;" Expression                                 // Case 88
//	|       Expression "\\[VerticalBar]" Expression                    // Case 89
//	|       Expression "\\[NotVerticalBar]" Expression                 // Case 90
//	|       Expression "\\[DoubleVerticalBar]" Expression              // Case 91
//	|       Expression "\\[NotDoubleVerticalBar]" Expression           // Case 92
//	|       Expression "\\[Element]" Expression                        // Case 93
//	|       Expression "\\[NotElement]" Expression                     // Case 94
//	|       Expression "\\[Subset]" Expression                         // Case 95
//	|       Expression "\\[Superset]" Expression                       // Case 96
//	|       Expression "\\[Nand]" Expression                           // Case 97
//	|       Expression "\\[Xor]" Expression                            // Case 98
//	|       Expression "\\[Xnor]" Expression                           // Case 99
//	|       Expression "\\[Nor]" Expression                            // Case 100
//	|       Expression "\\[Equivalent]" Expression                     // Case 101
//	|       Expression "\\[Implies]" Expression                        // Case 102
//	|       Expression "\\[RightTee]" Expression                       // Case 103
//	|       Expression "\\[DoubleRightTee]" Expression                 // Case 104
//	|       Expression "\\[LeftTee]" Expression                        // Case 105
//	|       Expression "\\[DoubleLeftTee]" Expression                  // Case 106
//	|       Expression "\\[UpTee]" Expression                          // Case 107
//	|       Expression "\\[DownTee]" Expression                        // Case 108
//	|       Expression "\\[SuchThat]" Expression                       // Case 109
//	|       Expression ".."                                            // Case 110
//	|       Expression "..."                                           // Case 111
//	|       Expression "~~" Expression                                 // Case 112
//	|       Expression "+=" Expression                                 // Case 113
//	|       Expression "-=" Expression                                 // Case 114
//	|       Expression "*=" Expression                                 // Case 115
//	|       Expression "/=" Expression                                 // Case 116
//	|       Expression "\\[VerticalSeparator]" Expression              // Case 117
//	|       Expression "\\[Therefore]" Expression                      // Case 118
//	|       Expression "\\[Because]" Expression                        // Case 119
//	|       Expression "^=" Expression                                 // Case 120
//	|       Expression "^:=" Expression                                // Case 121
//	|       Expression "/:" Expression                                 // Case 122
//	|       Expression '=' '.'                                         // Case 123
//	|       Expression "\\[Function]" Expression                       // Case 124
//	|       Expression ">>" FileName                                   // Case 125
//	|       Expression ">>>" FileName                                  // Case 126
//	|       Expression "\\`" STRING                                    // Case 127
//	|       IDENT ':' Expression                                       // Case 128
type Expression struct {
	Case        int
	Expression  *Expression
	Expression2 *Expression
	Expression3 *Expression
	Factor      *Factor
	FileName    *FileName
	Term        *Term
	Token       Token
	Token2      Token
}

func (n *Expression) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Expression) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Expression) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 56, 58, 59, 60, 62, 63, 64, 65, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127:
		return n.Expression.Pos()
	case 55:
		return n.Factor.Pos()
	case 0, 1, 2, 3, 4, 5, 6, 7, 40, 57, 61, 66, 67, 68, 85, 86, 128:
		return n.Token.Pos()
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

// FileName represents data reduced by productions:
//
//	FileName:
//	        IDENT
//	|       STRING  // Case 1
type FileName struct {
	Case  int
	Token Token
}

func (n *FileName) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *FileName) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *FileName) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// Tag represents data reduced by productions:
//
//	Tag:
//	        IDENT
//	|       STRING  // Case 1
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
//	        "<<" FileName
//	|       '(' Expression ')'                // Case 1
//	|       '{' '}'                           // Case 2
//	|       '{' ExprList CommaOpt '}'         // Case 3
//	|       FLOAT                             // Case 4
//	|       IDENT                             // Case 5
//	|       IDENT "::" Tag                    // Case 6
//	|       IDENT "::" Tag "::" Tag           // Case 7
//	|       INT                               // Case 8
//	|       PATTERN                           // Case 9
//	|       SLOT                              // Case 10
//	|       STRING                            // Case 11
//	|       Term "[[" ExprList CommaOpt "]]"  // Case 12
//	|       Term '!'                          // Case 13
//	|       Term '!' '!'                      // Case 14
//	|       Term '&'                          // Case 15
//	|       Term '[' ']'                      // Case 16
//	|       Term '[' ExprList CommaOpt ']'    // Case 17
//	|       Term QUOTE                        // Case 18
//	|       OUT                               // Case 19
type Term struct {
	Case       int
	CommaOpt   *CommaOpt
	ExprList   *ExprList
	Expression *Expression
	FileName   *FileName
	Tag        *Tag
	Tag2       *Tag
	Term       *Term
	Token      Token
	Token2     Token
	Token3     Token
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
	case 12, 13, 14, 15, 16, 17, 18:
		return n.Term.Pos()
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 19:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// start represents data reduced by production:
//
//	start:
//	        Expression
type start struct {
	Expression *Expression
}

func (n *start) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *start) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *start) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Expression.Pos()
}

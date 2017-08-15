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
	/*yy:token "%%%d"   */	OUT		"out"
	/*yy:token "%c"     */	IDENT		"identifier"
	/*yy:token "%c_"    */	PATTERN		"pattern"
	/*yy:token "%d"     */	INT		"integer"
	/*yy:token "1.%d"   */	FLOAT		"real"
	/*yy:token "\"%c\"" */	STRING		"string"

%token
	ADD_TO				"+="
	AND				"&&"
	APPLY				"@@"
	APPLY_ALL			"@@@"
	BACKSLASH			"\\[Backslash]"
	BECAUSE				"\\[Because]"
	CAP				"\\[Cap]"
	CENTER_DOT			"\\[CenterDot]"
	CIRCLE_DOT			"\\[CircleDot]"
	CIRCLE_MINUS			"\\[CircleMinus]"
	CIRCLE_PLUS			"\\[CirclePlus]"
	CIRCLE_TIMES			"\\[CircleTimes]"
	COMPOSITION			"@*"
	CONDITION			"/;"
	CONJUGATE			"\\[Conjugate]"
	CONJUGATE_TRANSPOSE		"\\[ConjugateTranspose]"
	COPRODUCT			"\\[Coproduct]"
	CROSS				"\\[Cross]"
	CUP				"\\[Cup]"
	DEC				"--"
	DEL				"\\[Del]"
	DIAMOND				"\\[Diamond]"
	DIFFERENCE_DELTA		"\\[DifferenceDelta]"
	DIFFERENTIAL_D			"\\[DifferentialD]"
	DISCRETE_RATIO			"\\[DiscreteRatio]"
	DISCRETE_SHIFT			"\\[DiscreteShift]"
	DIVIDE				"\\[Divide]"
	DIVIDE2				"\\/"
	DIVIDE_BY			"/="
	DOUBLE_LEFT_TEE			"\\[DoubleLeftTee]"
	DOUBLE_RIGHT_TEE		"\\[DoubleRightTee]"
	DOUBLE_VERTICAL_BAR		"\\[DoubleVerticalBar]"
	DOWN_TEE			"\\[DownTee]"
	ELEMENT				"\\[Element]"
	EQUAL				"=="
	EQUIVALENT			"\\[Equivalent]"
	FORM_BOX			"\\`"
	FOR_ALL				"\\[ForAll]"
	FUNCTION			"\\[Function]"
	GEQ				">="
	GET				"<<"
	HERMITIAN_CONJUGATE		"\\[HermitianConjugate]"
	IGNORE				// internal use only
	IMPLIES				"\\[Implies]"
	INC				"++"
	INTEGRATE			"\\[Integrate]"
	INTERSECTION			"\\[Intersection]"
	LEFT_TEE			"\\[LeftTee]"
	LEQ				"<="
	LPART				"[["
	MAP				"/@"
	MAP_ALL				"//@"
	MESSAGE_NAME			"::"
	MINUS_PLUS			"\\[MinusPlus]"
	NAND				"\\[Nand]"
	NON_COMMUTATIVE_MULTIPLY	"**"
	NOR				"\\[Nor]"
	NOT_DOUBLE_VERTICAL_BAR		"\\[NotDoubleVerticalBar]"
	NOT_ELEMENT			"\\[NotElement]"
	NOT_VERTICAL_BAR		"\\[NotVerticalBar]"
	OR				"||"
	OVERSCRIPT			"\\&"
	PARTIAL_D			"\\[PartialD]"
	PLUS_MINUS			"\\[PlusMinus]"
	POSTFIX				"//"
	POWER_SUBSCRIPT1		"\\^"
	POWER_SUBSCRIPT2		"\\%"
	PRODUCT				"\\[Product]"
	PUT				">>"
	PUT_APPEND			">>>"
	QUOTE				"'"
	REPEATED			".."
	REPEATED_NULL			"..."
	REPLACEALL			"/."
	REPLACEREP			"//."
	RIGHT_COMPOSITION		"/*"
	RIGHT_TEE			"\\[RightTee]"
	RPART				"]]"
	RULE				"->"
	RULEDELAYED			":>"
	SAME				"==="
	SET_DELAYED			":="
	SMALL_CIRCLE			"\\[SmallCircle]"
	SPAN				";;"
	SQRT				"\\[Sqrt]"
	SQRT2				"\\@"
	SQUARE				"\\[Square]"
	STAR				"\\[Star]"
	STRINGJOIN			"<>"
	STRING_EXPRESSION		"~~"
	SUBSCRIPT			"\\_"
	SUBSET				"\\[Subset]"
	SUBTRACT_FROM			"-="
	SUCH_THAT			"\\[SuchThat]"
	SUM				"\\[Sum]"
	SUPERSET			"\\[Superset]"
	TAG_SET				"/:"
	THEREFORE			"\\[Therefore]"
	TIMES_BY			"*="
	TRANSPOSE			"\\[Transpose]"
	UNDERSCRIPT			"\\+"
	UNEQUAL				"!="
	UNION				"\\[Union]"
	UNSAME				"=!="
	UP_SET				"^="
	UP_SET_DELAYED			"^:="
	UP_TEE				"\\[UpTee]"
	VEE				"\\[Vee]"
	VERTICAL_BAR			"\\[VerticalBar]"
	VERTICAL_SEPARATOR		"\\[VerticalSeparator]"
	VERTICAL_TILDE			"\\[VerticalTilde]"
	WEDGE				"\\[Wedge]"
	XNOR				"\\[Xnor]"
	XOR				"\\[Xor]"

%type	<Node>
	CommaOpt	"optional comma"
	ExprList	"expression list"
	Expression	"expression"
	Factor		"factor"
	FileName	"file name"
	Tag		"tag"
	Term		"term"
	start		"valid input"

%left		FORM_BOX
%left		';'					// CompoundExpression
%right		PUT PUT_APPEND
%right		'=' SET_DELAYED UP_SET UP_SET_DELAYED TAG_SET FUNCTION
%left		BECAUSE
%right		THEREFORE
%left		VERTICAL_SEPARATOR
%left		POSTFIX
%left		':'					// Colon
%precedence	'&'					// Function
%right		ADD_TO SUBTRACT_FROM TIMES_BY DIVIDE_BY
%left		REPLACEALL REPLACEREP
%left		RULE RULEDELAYED
%left		CONDITION
%left		STRING_EXPRESSION
%left		PATTERN
%left		'|'					// Alternatives
%nonassoc	REPEATED REPEATED_NULL
%right		SUCH_THAT
%left		LEFT_TEE DOUBLE_LEFT_TEE UP_TEE DOWN_TEE
%right		RIGHT_TEE DOUBLE_RIGHT_TEE
%right		IMPLIES
%left		EQUIVALENT
%left		OR NOR
%left		XOR XNOR
%left		AND NAND
%right		'!'					// Not
%left		FOR_ALL EXISTS NOT_EXISTS
%left		ELEMENT NOT_ELEMENT SUBSET SUPERSET
%left		SAME UNSAME
%left		EQUAL UNEQUAL '>' GEQ '<' LEQ VERTICAL_BAR NOT_VERTICAL_BAR DOUBLE_VERTICAL_BAR NOT_DOUBLE_VERTICAL_BAR
%left		SPAN
%left		UNION
%left		INTERSECTION
%left		'+' '-' PLUS_MINUS MINUS_PLUS
%right		SUM
%left		CIRCLE_PLUS CIRCLE_MINUS
%left		CAP CUP
%left		COPRODUCT
%left		VERTICAL_TILDE
%right		PRODUCT
%left		STAR
%left		'*'
%left		CENTER_DOT
%left		CIRCLE_TIMES
%left		VEE
%left		WEDGE
%left		DIAMOND
%left		BACKSLASH
%left		'/'
%precedence	UNARY_MINUS UNARY_PLUS UNARY_PLUS_MINUS UNARY_MINUS_PLUS
%left		'.'					// Dot
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
%precedence	PRE_INC					// PreIncrement, PreDecrement
%nonassoc	INC DEC					// Increment, Decrement
%left		'[' ']' LPART RPART			// expr, Part
%left		'?'					// PatternTest
%right		SUBSCRIPT
%right		OVERSCRIPT UNDERSCRIPT
%nonassoc	GET
%left		MESSAGE_NAME

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
|	PATTERN ':' Term
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
|	"\\[PlusMinus]" Expression %prec UNARY_PLUS_MINUS
|	"\\[MinusPlus]" Expression %prec UNARY_MINUS_PLUS
|	Expression "\\[Backslash]" Expression
|	Expression "!=" Expression
|	Expression "\\[Diamond]" Expression
|	Expression "\\[Wedge]" Expression
|	Expression "\\[Vee]" Expression
|	Expression "\\[CircleTimes]" Expression
|	Expression "\\[CenterDot]" Expression
|	Expression "\\[Star]" Expression
|	Expression "\\[VerticalTilde]" Expression
|	Expression "\\[Coproduct]" Expression
|	Expression "\\[Cap]" Expression
|	Expression "\\[Cup]" Expression
|	Expression "\\[CirclePlus]" Expression
|	Expression "\\[CircleMinus]" Expression
|	Expression "\\[Intersection]" Expression
|	Expression "\\[Union]" Expression
|	";;"
|	";;" Expression
|	Expression ";;"
|	Expression ";;" Expression
|	Expression "\\[VerticalBar]" Expression
|	Expression "\\[NotVerticalBar]" Expression
|	Expression "\\[DoubleVerticalBar]" Expression
|	Expression "\\[NotDoubleVerticalBar]" Expression
|	Expression "\\[Element]" Expression
|	Expression "\\[NotElement]" Expression
|	Expression "\\[Subset]" Expression
|	Expression "\\[Superset]" Expression
|	Expression "\\[Nand]" Expression
|	Expression "\\[Xor]" Expression
|	Expression "\\[Xnor]" Expression
|	Expression "\\[Nor]" Expression
|	Expression "\\[Equivalent]" Expression
|	Expression "\\[Implies]" Expression
|	Expression "\\[RightTee]" Expression
|	Expression "\\[DoubleRightTee]" Expression
|	Expression "\\[LeftTee]" Expression
|	Expression "\\[DoubleLeftTee]" Expression
|	Expression "\\[UpTee]" Expression
|	Expression "\\[DownTee]" Expression
|	Expression "\\[SuchThat]" Expression
|	Expression ".."
|	Expression "..."
|	Expression "~~" Expression
|	Expression "+=" Expression
|	Expression "-=" Expression
|	Expression "*=" Expression
|	Expression "/=" Expression
|	Expression "\\[VerticalSeparator]" Expression
|	Expression "\\[Therefore]" Expression
|	Expression "\\[Because]" Expression
|	Expression "^=" Expression
|	Expression "^:=" Expression
//yy:example "a/:b=c"
|	Expression "/:" Expression
	{
		switch lhs.Expression2.Case {
		case
			19, // Expression ":=" Expression                                 // Case 19
			44: // Expression '=' Expression                                  // Case 44

			// ok
		default:
			lx.errPos(lhs.Expression2.Pos(), "expected 'Expression = Expression' or 'Expression := Expression'")
		}
	}
|	Expression '=' '.'
|	Expression "\\[Function]" Expression
|	Expression ">>" FileName
|	Expression ">>>" FileName
|	Expression "\\`" STRING
|	IDENT ':' Expression

Term:
	"<<" FileName
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
|	OUT

Factor:
	Term
|	Term Factor

ExprList:
	Expression
|	ExprList ',' Expression

CommaOpt:
	/* empty */ {}
|	','

FileName:
	IDENT
|	STRING

Tag:
	IDENT
|	STRING

%%

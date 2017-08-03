// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

/*

Operator Precedence

operator form
full form
grouping
forms representing numbers (see Numbers)	⊲
forms representing symbols (see Symbol Names and Contexts)	⊲
forms representing character strings (see Character Strings)	⊲
e11	e12	…
e21	e22	…
…		
{{e11,e12,…},{e21,e22,…},…}		⊲

e11	e12
e21	e22
…	
Piecewise[{{e11,e12},{e21,e22},…}]		⊲
expr::string	MessageName[expr,"string"]		⊲
expr::string1::string2	MessageName[expr,"string1","string2"]		⊲
forms containing # (see additional input forms)	⊲
forms containing % (see additional input forms)	⊲
forms containing _ (see additional input forms)	⊲
<<filename	Get["filename"]		⊲
	Overscript[expr1,expr2]		
expr1\&expr2	Overscript[expr1,expr2]	e\&(e\&e)	
	Underscript[expr1,expr2]		
expr1\+expr2	Underscript[expr1,expr2]	e\+(e\+e)	
	Underoverscript[expr1,expr2,expr3]		
expr1\+expr2\%expr3	Underoverscript[expr1,expr2,expr3]		
expr1\&expr2\%expr3	Underoverscript[expr1,expr3,expr2]		
expr1expr2	Subscript[expr1,expr2]	e(ee)	
expr1\_expr2	Subscript[expr1,expr2]	e\_(e\_e)	
expr1\_expr2\%expr3	Power[Subscript[expr1,expr2],expr3]		⊲
\!boxes	(interpreted version of boxes)		
expr1?expr2	PatternTest[expr1,expr2]		⊲
expr1[expr2,…]	expr1[expr2,…]	(e[e])[e]	⊲
expr1[[expr2,…]]	Part[expr1,expr2,…]	(e[[e]])[[e]]	⊲
expr1〚expr2,…〛	Part[expr1,expr2,…]	(e〚e〛)〚e〛	⊲
expr1〚expr2〛	Part[expr1,expr2,…]	(e〚e〛)〚e〛	⊲
\*expr	(boxes constructed from expr)		
expr++	Increment[expr]		⊲
expr--	Decrement[expr]		⊲
++expr	PreIncrement[expr]		⊲
--expr	PreDecrement[expr]		⊲
expr1@*expr2	Composition[expr1,expr2]	e@*e@*e	⊲
expr1/*expr2	RightComposition[expr1,expr2]	e/*e/*e	⊲
expr1@expr2	expr1[expr2]	e@(e@e)	⊲
expr1 expr2	(invisible application, input as expr1 Esc@Esc expr2)	⊲
expr1[expr2]		
expr1~expr2~expr3	expr2[expr1,expr3]	(e~e~e)~e~e	⊲
expr1/@expr2	Map[expr1,expr2]	e/@(e/@e)	⊲
expr1//@expr2	MapAll[expr1,expr2]	e//@(e//@e)	⊲
expr1@@expr2	Apply[expr1,expr2]	e@@(e@@e)	⊲
expr1@@@expr2	Apply[expr1,expr2,{1}]	e@@@(e@@@e)	⊲
expr!	Factorial[expr]		⊲
expr!!	Factorial2[expr]		⊲
expr	Conjugate[expr]		⊲
expr	Transpose[expr]		⊲
expr	ConjugateTranspose[expr]		⊲
expr	ConjugateTranspose[expr]		⊲
expr'	Derivative[1][expr]		⊲
expr''…' (n times)	Derivative[n][expr]		⊲
expr1<>expr2<>expr3	StringJoin[expr1,expr2,expr3]	e<>e<>e	⊲
expr1^expr2	Power[expr1,expr2]	e^(e^e)	⊲
expr1expr2	Power[expr1,expr2]	e(ee)	⊲
	Power[Subscript[expr1,expr2],expr3]		⊲
expr1\^expr2\%expr3	Power[Subscript[expr1,expr3],expr2]		⊲
vertical arrow and vector operators
	Sqrt[expr]		⊲
\@ expr	Sqrt[expr]	\@(\@ e)	⊲
\@ expr\%n	Power[expr,1/n]		⊲
∫ expr1 expr2	Integrate[expr1,expr2]	∫ (∫ e e) e	⊲
e3e4	Integrate[e3,{e4,e1,e2}]	∫ (∫ e e) e	⊲
other integration operators
∂expr1expr2	D[expr2,expr1]	∂e(∂ee)	⊲
∇ expr	Del[expr]	∇(∇e)	
expr1expr2	DiscreteShift[expr2,expr1]	e(ee)	⊲
expr1expr2	DiscreteRatio[expr2,expr1]	e(ee)	⊲
expr1expr2	DifferenceDelta[expr2,expr1]	e(ee)	⊲
 expr	Square[expr]	( e)	
expr1∘ expr2∘ expr3	SmallCircle[expr1,expr2,expr3]	e∘ e∘ e	
expr1⊙ expr2⊙ expr3	CircleDot[expr1,expr2,expr3]	e ⊙ e ⊙ e	
expr1**expr2**expr3	NonCommutativeMultiply[expr1,expr2,expr3]	e**e**e	
expr1expr2expr3	Cross[expr1,expr2,expr3]	eee	⊲
expr1.expr2.expr3	Dot[expr1,expr2,expr3]	e.e.e	⊲
-expr	Times[-1,expr]		⊲
+expr	expr		⊲
±expr	PlusMinus[expr]		
∓expr	MinusPlus[expr]		
expr1/expr2	expr1(expr2)^-1	(e/e)/e	⊲
expr1÷expr2	Divide[expr1,expr2]	(e÷e)÷e	⊲
expr1\/expr2	Divide[expr1,expr2]	(e\/e)\/e	⊲
expr1∖expr2∖expr3	Backslash[expr1,expr2,expr3]	e∖e∖e	
expr1⋄expr2⋄expr3	Diamond[expr1,expr2,expr3]	e⋄e⋄e	
expr1⋀expr2⋀expr3	Wedge[expr1,expr2,expr3]	e⋀e⋀e	
expr1⋁expr2⋁expr3	Vee[expr1,expr2,expr3]	e⋁e⋁e	
expr1⊗expr2⊗expr3	CircleTimes[expr1,expr2,expr3]	e⊗e⊗e	
expr1·expr2·expr3	CenterDot[expr1,expr2,expr3]	e·e·e	
expr1 expr2 expr3	Times[expr1,expr2,expr3]	e e e	⊲
expr1*expr2*expr3	Times[expr1,expr2,expr3]	e*e*e	⊲
expr1×expr2×expr3	Times[expr1,expr2,expr3]	e×e×e	⊲
expr1⋆expr2⋆expr3	Star[expr1,expr2,expr3]	e⋆e⋆e	
e4	Product[e4,{e1,e2,e3}]	∏(∏ e)	⊲
expr1≀expr2≀expr3	VerticalTilde[expr1,expr2,expr3]	e≀e≀e	
expr1∐expr2∐expr3	Coproduct[expr1,expr2,expr3]	e∐e∐e	
expr1⌢expr2⌢expr3	Cap[expr1,expr2,expr3]	e⌢e⌢e	
expr1⌣expr2⌣expr3	Cup[expr1,expr2,expr3]	e⌣e⌣e	
expr1⊕ expr2⊕ expr3	CirclePlus[expr1,expr2,expr3]	e⊕e⊕e	
expr1⊖ expr2	CircleMinus[expr1,expr2]	(e ⊖ e)⊖ e	
e4	Sum[e4,{e1,e2,e3}]	∑(∑ e)	⊲
expr1+expr2+expr3	Plus[expr1,expr2,expr3]	e+e+e	⊲
expr1-expr2	expr1+(-1expr2)	(e-e)-e	⊲
expr1±expr2	PlusMinus[expr1,expr2]	(e±e)±e	
expr1∓expr2	MinusPlus[expr1,expr2]	(e∓e)∓e	
expr1⋂expr2	Intersection[expr1,expr2]	e⋂e⋂e	⊲
other intersection operators
expr1⋃expr2	Union[expr1,expr2]	e⋃e⋃e	⊲
other union operators
i;;j;;k	Span[i,j,k]	e;;e;;e	⊲
expr1==expr2	Equal[expr1,expr2]	e==e==e	⊲
expr1==expr2	Equal[expr1,expr2]	e==e==e	⊲
expr1expr2	Equal[expr1,expr2]	eee	⊲
expr1!= expr2	Unequal[expr1,expr2]	e!=e!=e	⊲
expr1!=expr2	Unequal[expr1,expr2]	e!=e!=e	⊲
other equality and similarity operators
expr1>expr2	Greater[expr1,expr2]	e>e>e	⊲
expr1>=expr2	GreaterEqual[expr1,expr2]	e>=e>=e	⊲
expr1≥expr2	GreaterEqual[expr1,expr2]	e≥e≥e	⊲
expr1⩾expr2	GreaterEqual[expr1,expr2]	e⩾e⩾e	⊲
expr1<expr2	Less[expr1,expr2]	e<e<e	⊲
expr1<=expr2	LessEqual[expr1,expr2]	e<=e<=e	⊲
expr1≤expr2	LessEqual[expr1,expr2]	e≤e≤e	⊲
expr1⩽expr2	LessEqual[expr1,expr2]	e⩽e⩽e	⊲
other ordering operators
expr1expr2	VerticalBar[expr1,expr2]	eee	
expr1expr2	NotVerticalBar[expr1,expr2]	eee	
expr1∥expr2	DoubleVerticalBar[expr1,expr2]	e∥e∥e	
expr1∦expr2	NotDoubleVerticalBar[expr1,expr2]	e∦e∦e	
horizontal arrow and vector operators
diagonal arrow operators			
expr1===expr2	SameQ[expr1,expr2]	e===e===e	⊲
expr1=!=expr2	UnsameQ[expr1,expr2]	e=!=e=!=e	⊲
expr1∈expr2	Element[expr1,expr2]	e∈e∈e	⊲
expr1∉expr2	NotElement[expr1,expr2]	e∉e∉e	⊲
expr1⊂expr2	Subset[expr1,expr2]	e⊂e⊂e	
expr1⊃expr2	Superset[expr1,expr2]	e⊃e⊃e	
other set relation operators	
∀expr1expr2	ForAll[expr1,expr2]	∀e(∀ee)	⊲
∃expr1expr2	Exists[expr1,expr2]	∃e(∃ee)	⊲
∄expr1expr2	NotExists[expr1,expr2]	∄e(∄ee)	
!expr	Not[expr]	!(!e)	⊲
¬expr	Not[expr]	¬(¬e)	⊲
expr1&&expr2&&expr3	And[expr1,expr2,expr3]	e&&e&&e	⊲
expr1∧expr2∧expr3	And[expr1,expr2,expr3]	e∧e∧e	⊲
expr1⊼expr2⊼expr3	Nand[expr1,expr2,expr3]	e⊼e⊼e	⊲
expr1⊻expr2⊻expr3	Xor[expr1,expr2,expr3]	e⊻e⊻e	⊲
expr1expr2expr3	Xnor[expr1,expr2,expr3]	eee	⊲
expr1||expr2||expr3	Or[expr1,expr2,expr3]	e||e||e	⊲
expr1∨expr2∨expr3	Or[expr1,expr2,expr3]	e∨e∨e	⊲
expr1⊽expr2⊽expr3	Nor[expr1,expr2,expr3]	e⊽e⊽e	⊲
expr1⧦expr2⧦expr3	Equivalent[expr1,expr2,expr3]	e⧦e⧦e	⊲
expr1expr2	Implies[expr1,expr2]	e(ee)	⊲
expr1⥰expr2	Implies[expr1,expr2]	e⥰e⥰e	⊲
expr1⊢expr2	RightTee[expr1,expr2]	e⊢(e⊢e)	
expr1⊨expr2	DoubleRightTee[expr1,expr2]	e⊨(e⊨e)	
expr1⊣expr2	LeftTee[expr1,expr2]	(e⊣e)⊣e	
expr1⫤expr2	DoubleLeftTee[expr1,expr2]	(e⫤e)⫤e	
expr1⊥expr2	UpTee[expr1,expr2]	(e⊥e)⊥e	
expr1⊤expr2	DownTee[expr1,expr2]	(e⊤e)⊤e	
expr1∍expr2	SuchThat[expr1,expr2]	e∍(e∍e)	
expr..	Repeated[expr]		⊲
expr...	RepeatedNull[expr]		⊲
expr1|expr2	Alternatives[expr1,expr2]	e|e|e	⊲
symb:expr	Pattern[symb,expr]		⊲
patt:expr	Optional[patt,expr]		⊲
expr1~~expr2~~expr3	StringExpression[expr1,expr2,expr3]	e~~e~~e	⊲
expr1/;expr2	Condition[expr1,expr2]	(e/;e)/;e	⊲
expr1->expr2	Rule[expr1,expr2]	e->(e->e)	⊲
expr1expr2	Rule[expr1,expr2]	e(ee)	⊲
expr1:>expr2	RuleDelayed[expr1,expr2]	e:>(e:>e)	⊲
expr1 expr2	RuleDelayed[expr1,expr2]	e(ee)	⊲
expr1/.expr2	ReplaceAll[expr1,expr2]	(e/.e)/.e	⊲
expr1//.expr2	ReplaceRepeated[expr1,expr2]	(e//.e)//.e	⊲
expr1+=expr2	AddTo[expr1,expr2]	e+=(e+=e)	⊲
expr1-=expr2	SubtractFrom[expr1,expr2]	e-=(e-=e)	⊲
expr1*=expr2	TimesBy[expr1,expr2]	e*=(e*=e)	⊲
expr1/=expr2	DivideBy[expr1,expr2]	e/=(e/=e)	⊲
expr&	Function[expr]		⊲
expr1∶expr2	Colon[expr1:expr2]	e∶e∶e	
expr1//expr2	expr2[expr1]	(e//e)//e	
expr1expr2	VerticalSeparator[expr1,expr2]	eee	
expr1∴expr2	Therefore[expr1,expr2]	e∴(e∴e)	
expr1∵expr2	Because[expr1,expr2]	(e∵e)∵e	
expr1=expr2	Set[expr1,expr2]	e=(e=e)	⊲
expr1:=expr2	SetDelayed[expr1,expr2]	e:=(e:=e)	⊲
expr1^=expr2	UpSet[expr1,expr2]	e^=(e^=e)	⊲
expr1^:=expr2	UpSetDelayed[expr1,expr2]	e^:=(e^:=e)	⊲
symb/:expr1=expr2	TagSet[symb,expr1,expr2]		⊲
symb/:expr1:=expr2	TagSetDelayed[symb,expr1,expr2]		⊲
expr=.	Unset[expr]		⊲
symb/:expr=.	TagUnset[symb,expr]		⊲
expr1expr2	Function[{expr1},expr2]	e(ee)	⊲
expr>>filename	Put[expr,"filename"]		⊲
expr>>>filename	PutAppend[expr,"filename"]		⊲
expr1;expr2;expr3	CompoundExpression[expr1,expr2,expr3]		⊲
expr1;expr2;	CompoundExpression[expr1,expr2,Null]		⊲
expr1\`expr2	FormBox[expr2,expr1]	e\`(e\`e)	⊲
Operator input forms, in order of decreasing precedence. Operators of equal precedence are grouped together.

*/

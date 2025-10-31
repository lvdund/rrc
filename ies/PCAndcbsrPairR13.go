package ies

// P-C-AndCBSR-Pair-r13 ::= SEQUENCE OF P-C-AndCBSR-r13
// SIZE (1..2)
type PCAndcbsrPairR13 struct {
	Value []PCAndcbsrR13 `lb:1,ub:2`
}

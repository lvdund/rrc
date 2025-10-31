package ies

// P-C-AndCBSR-Pair-r15 ::= SEQUENCE OF P-C-AndCBSR-r15
// SIZE (1..2)
type PCAndcbsrPairR15 struct {
	Value []PCAndcbsrR15 `lb:1,ub:2`
}

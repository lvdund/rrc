package ies

// P-C-AndCBSR-Pair-r13a ::= SEQUENCE OF P-C-AndCBSR-r11
// SIZE (1..2)
type PCAndcbsrPairR13a struct {
	Value []PCAndcbsrR11 `lb:1,ub:2`
}

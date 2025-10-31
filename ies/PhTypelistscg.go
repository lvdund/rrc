package ies

// PH-TypeListSCG ::= SEQUENCE OF PH-InfoSCG
// SIZE (1..maxNrofServingCells)
type PhTypelistscg struct {
	Value []PhInfoscg `lb:1,ub:maxNrofServingCells`
}

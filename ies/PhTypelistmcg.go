package ies

// PH-TypeListMCG ::= SEQUENCE OF PH-InfoMCG
// SIZE (1..maxNrofServingCells)
type PhTypelistmcg struct {
	Value []PhInfomcg `lb:1,ub:maxNrofServingCells`
}

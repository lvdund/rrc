package ies

// IntraBandCC-Combination-r17 ::= SEQUENCE OF CC-State-r17
// SIZE (1.. maxNrofServingCells)
type IntrabandccCombinationR17 struct {
	Value []CcStateR17 `lb:1,ub:maxNrofServingCells`
}

package ies

// FR-InfoList ::= SEQUENCE OF FR-Info
// SIZE (1..maxNrofServingCells-1)
type FrInfolist struct {
	Value []FrInfo `lb:1,ub:maxNrofServingCells1`
}

package ies

// PerCC-GapIndicationList-r14 ::= SEQUENCE OF PerCC-GapIndication-r14
// SIZE (1..maxServCell-r13)
type PerccGapindicationlistR14 struct {
	Value []PerccGapindicationR14 `lb:1,ub:maxServCellR13`
}

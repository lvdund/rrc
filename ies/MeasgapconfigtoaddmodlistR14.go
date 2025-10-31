package ies

// MeasGapConfigToAddModList-r14 ::= SEQUENCE OF MeasGapConfigPerCC-r14
// SIZE (1..maxServCell-r13)
type MeasgapconfigtoaddmodlistR14 struct {
	Value []MeasgapconfigperccR14 `lb:1,ub:maxServCellR13`
}

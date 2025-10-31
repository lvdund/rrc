package ies

// MeasGapConfigToRemoveList-r14 ::= SEQUENCE OF ServCellIndex-r13
// SIZE (1..maxServCell-r13)
type MeasgapconfigtoremovelistR14 struct {
	Value []ServcellindexR13 `lb:1,ub:maxServCellR13`
}

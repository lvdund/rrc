package ies

// ValidityCellList-r16 ::= SEQUENCE OF PhysCellIdRange
// SIZE (1.. maxCellMeasIdle-r15)
type ValiditycelllistR16 struct {
	Value []Physcellidrange `lb:1,ub:maxCellMeasIdleR15`
}

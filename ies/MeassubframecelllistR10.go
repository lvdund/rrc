package ies

// MeasSubframeCellList-r10 ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellMeas)
type MeassubframecelllistR10 struct {
	Value []Physcellidrange `lb:1,ub:maxCellMeas`
}

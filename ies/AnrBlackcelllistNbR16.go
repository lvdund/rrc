package ies

// ANR-BlackCellList-NB-r16 ::= SEQUENCE OF PhysCellId
// SIZE (1..maxCellBlack)
type AnrBlackcelllistNbR16 struct {
	Value []Physcellid `lb:1,ub:maxCellBlack`
}

package ies

// InterFreqBlackCellList-NB-r13 ::= SEQUENCE OF PhysCellId
// SIZE (1..maxCellBlack)
type InterfreqblackcelllistNbR13 struct {
	Value []Physcellid `lb:1,ub:maxCellBlack`
}

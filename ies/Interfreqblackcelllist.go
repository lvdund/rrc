package ies

// InterFreqBlackCellList ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellBlack)
type Interfreqblackcelllist struct {
	Value []Physcellidrange `lb:1,ub:maxCellBlack`
}

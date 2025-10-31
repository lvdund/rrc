package ies

// IntraFreqBlackCellList ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellBlack)
type Intrafreqblackcelllist struct {
	Value []Physcellidrange `lb:1,ub:maxCellBlack`
}

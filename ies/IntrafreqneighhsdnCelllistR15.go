package ies

// IntraFreqNeighHSDN-CellList-r15 ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellIntra)
type IntrafreqneighhsdnCelllistR15 struct {
	Value []Physcellidrange `lb:1,ub:maxCellIntra`
}

package ies

// InterFreqNeighHSDN-CellList-r15 ::= SEQUENCE OF PhysCellIdRange
// SIZE (1..maxCellInter)
type InterfreqneighhsdnCelllistR15 struct {
	Value []Physcellidrange `lb:1,ub:maxCellInter`
}

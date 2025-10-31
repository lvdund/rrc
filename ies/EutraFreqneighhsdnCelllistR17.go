package ies

// EUTRA-FreqNeighHSDN-CellList-r17 ::= SEQUENCE OF EUTRA-PhysCellIdRange
// SIZE (1..maxCellEUTRA)
type EutraFreqneighhsdnCelllistR17 struct {
	Value []EutraPhyscellidrange `lb:1,ub:maxCellEUTRA`
}

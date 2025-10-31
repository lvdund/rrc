package ies

// InterFreqNeighHSDN-CellList-r17 ::= SEQUENCE OF PCI-Range
// SIZE (1..maxCellInter)
type InterfreqneighhsdnCelllistR17 struct {
	Value []PciRange `lb:1,ub:maxCellInter`
}

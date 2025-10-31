package ies

// IntraFreqNeighHSDN-CellList-r17 ::= SEQUENCE OF PCI-Range
// SIZE (1..maxCellIntra)
type IntrafreqneighhsdnCelllistR17 struct {
	Value []PciRange `lb:1,ub:maxCellIntra`
}

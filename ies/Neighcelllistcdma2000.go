package ies

// NeighCellListCDMA2000 ::= SEQUENCE OF NeighCellCDMA2000
// SIZE (1..16)
type Neighcelllistcdma2000 struct {
	Value []Neighcellcdma2000 `lb:1,ub:16`
}

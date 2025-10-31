package ies

// NeighCellsPerBandclassListCDMA2000 ::= SEQUENCE OF NeighCellsPerBandclassCDMA2000
// SIZE (1..16)
type Neighcellsperbandclasslistcdma2000 struct {
	Value []Neighcellsperbandclasscdma2000 `lb:1,ub:16`
}

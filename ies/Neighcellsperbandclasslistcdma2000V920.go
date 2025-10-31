package ies

// NeighCellsPerBandclassListCDMA2000-v920 ::= SEQUENCE OF NeighCellsPerBandclassCDMA2000-v920
// SIZE (1..16)
type Neighcellsperbandclasslistcdma2000V920 struct {
	Value []Neighcellsperbandclasscdma2000V920 `lb:1,ub:16`
}

package ies

// InterFreqNeighCellList-NB-r13 ::= SEQUENCE OF PhysCellId
// SIZE (1..maxCellInter)
type InterfreqneighcelllistNbR13 struct {
	Value []Physcellid `lb:1,ub:maxCellInter`
}

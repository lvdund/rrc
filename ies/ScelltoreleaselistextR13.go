package ies

// SCellToReleaseListExt-r13 ::= SEQUENCE OF SCellIndex-r13
// SIZE (1..maxSCell-r13)
type ScelltoreleaselistextR13 struct {
	Value []ScellindexR13 `lb:1,ub:maxSCellR13`
}

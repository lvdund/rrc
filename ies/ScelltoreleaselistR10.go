package ies

// SCellToReleaseList-r10 ::= SEQUENCE OF SCellIndex-r10
// SIZE (1..maxSCell-r10)
type ScelltoreleaselistR10 struct {
	Value []ScellindexR10 `lb:1,ub:maxSCellR10`
}

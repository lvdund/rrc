package ies

// SCellToAddModList-r10 ::= SEQUENCE OF SCellToAddMod-r10
// SIZE (1..maxSCell-r10)
type ScelltoaddmodlistR10 struct {
	Value []ScelltoaddmodR10 `lb:1,ub:maxSCellR10`
}

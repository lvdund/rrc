package ies

// SCellToAddModList-v13c0 ::= SEQUENCE OF SCellToAddMod-v13c0
// SIZE (1..maxSCell-r10)
type ScelltoaddmodlistV13c0 struct {
	Value []ScelltoaddmodV13c0 `lb:1,ub:maxSCellR10`
}

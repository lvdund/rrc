package ies

// SCellToAddModList-v10l0 ::= SEQUENCE OF SCellToAddMod-v10l0
// SIZE (1..maxSCell-r10)
type ScelltoaddmodlistV10l0 struct {
	Value []ScelltoaddmodV10l0 `lb:1,ub:maxSCellR10`
}

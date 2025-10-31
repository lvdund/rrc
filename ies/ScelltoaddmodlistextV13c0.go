package ies

// SCellToAddModListExt-v13c0 ::= SEQUENCE OF SCellToAddMod-v13c0
// SIZE (1..maxSCell-r13)
type ScelltoaddmodlistextV13c0 struct {
	Value []ScelltoaddmodV13c0 `lb:1,ub:maxSCellR13`
}

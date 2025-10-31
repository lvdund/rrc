package ies

// SCellToAddModListExt-r13 ::= SEQUENCE OF SCellToAddModExt-r13
// SIZE (1..maxSCell-r13)
type ScelltoaddmodlistextR13 struct {
	Value []ScelltoaddmodextR13 `lb:1,ub:maxSCellR13`
}

package ies

// SCellToAddModListExt-v1370 ::= SEQUENCE OF SCellToAddModExt-v1370
// SIZE (1..maxSCell-r13)
type ScelltoaddmodlistextV1370 struct {
	Value []ScelltoaddmodextV1370 `lb:1,ub:maxSCellR13`
}

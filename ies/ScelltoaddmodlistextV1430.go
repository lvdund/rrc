package ies

// SCellToAddModListExt-v1430 ::= SEQUENCE OF SCellToAddModExt-v1430
// SIZE (1..maxSCell-r13)
type ScelltoaddmodlistextV1430 struct {
	Value []ScelltoaddmodextV1430 `lb:1,ub:maxSCellR13`
}

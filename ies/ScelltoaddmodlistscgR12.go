package ies

// SCellToAddModListSCG-r12 ::= SEQUENCE OF Cell-ToAddMod-r12
// SIZE (1..maxSCell-r10)
type ScelltoaddmodlistscgR12 struct {
	Value []CellToaddmodR12 `lb:1,ub:maxSCellR10`
}

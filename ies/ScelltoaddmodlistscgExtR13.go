package ies

// SCellToAddModListSCG-Ext-r13 ::= SEQUENCE OF Cell-ToAddMod-r12
// SIZE (1..maxSCell-r13)
type ScelltoaddmodlistscgExtR13 struct {
	Value []CellToaddmodR12 `lb:1,ub:maxSCellR13`
}

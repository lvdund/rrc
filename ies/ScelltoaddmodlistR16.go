package ies

// SCellToAddModList-r16 ::= SEQUENCE OF SCellToAddMod-r16
// SIZE (1..maxSCell-r13)
type ScelltoaddmodlistR16 struct {
	Value []ScelltoaddmodR16 `lb:1,ub:maxSCellR13`
}

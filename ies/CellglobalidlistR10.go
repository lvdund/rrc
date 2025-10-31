package ies

// CellGlobalIdList-r10 ::= SEQUENCE OF CellGlobalIdEUTRA
// SIZE (1..32)
type CellglobalidlistR10 struct {
	Value []Cellglobalideutra `lb:1,ub:32`
}

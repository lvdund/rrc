package ies

import "rrc/utils"

// CellGlobalIdList-r10 ::= SEQUENCE OF CellGlobalIdEUTRA
// SIZE (1..32)
type CellglobalidlistR10 struct {
	Value []Cellglobalideutra
}

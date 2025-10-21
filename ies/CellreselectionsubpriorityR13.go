package ies

import "rrc/utils"

// CellReselectionSubPriority-r13 ::= ENUMERATED
type CellreselectionsubpriorityR13 struct {
	Value utils.ENUMERATED
}

const (
	CellreselectionsubpriorityR13Odot2 = 0
	CellreselectionsubpriorityR13Odot4 = 1
	CellreselectionsubpriorityR13Odot6 = 2
	CellreselectionsubpriorityR13Odot8 = 3
)

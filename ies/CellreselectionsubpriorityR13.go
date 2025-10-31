package ies

import "rrc/utils"

// CellReselectionSubPriority-r13 ::= ENUMERATED
type CellreselectionsubpriorityR13 struct {
	Value utils.ENUMERATED
}

const (
	CellreselectionsubpriorityR13EnumeratedNothing = iota
	CellreselectionsubpriorityR13EnumeratedOdot2
	CellreselectionsubpriorityR13EnumeratedOdot4
	CellreselectionsubpriorityR13EnumeratedOdot6
	CellreselectionsubpriorityR13EnumeratedOdot8
)

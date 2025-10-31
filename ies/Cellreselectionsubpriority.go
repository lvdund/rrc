package ies

import "rrc/utils"

// CellReselectionSubPriority ::= ENUMERATED
type Cellreselectionsubpriority struct {
	Value utils.ENUMERATED
}

const (
	CellreselectionsubpriorityEnumeratedNothing = iota
	CellreselectionsubpriorityEnumeratedOdot2
	CellreselectionsubpriorityEnumeratedOdot4
	CellreselectionsubpriorityEnumeratedOdot6
	CellreselectionsubpriorityEnumeratedOdot8
)

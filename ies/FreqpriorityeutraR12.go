package ies

import "rrc/utils"

// FreqPriorityEUTRA-r12 ::= SEQUENCE
type FreqpriorityeutraR12 struct {
	CarrierfreqR12             ArfcnValueeutraR9
	CellreselectionpriorityR12 Cellreselectionpriority
}

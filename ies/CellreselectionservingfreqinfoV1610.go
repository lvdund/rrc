package ies

import "rrc/utils"

// CellReselectionServingFreqInfo-v1610 ::= SEQUENCE
type CellreselectionservingfreqinfoV1610 struct {
	AltcellreselectionpriorityR16    *Cellreselectionpriority
	AltcellreselectionsubpriorityR16 *CellreselectionsubpriorityR13
}

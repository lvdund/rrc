package ies

import "rrc/utils"

// SL-SelectionWindowConfig-r16 ::= SEQUENCE
type SlSelectionwindowconfigR16 struct {
	SlPriorityR16        utils.INTEGER `lb:0,ub:8`
	SlSelectionwindowR16 SlSelectionwindowconfigR16SlSelectionwindowR16
}

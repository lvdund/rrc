package ies

import "rrc/utils"

// SL-PTRS-Config-r16 ::= SEQUENCE
// Extensible
type SlPtrsConfigR16 struct {
	SlPtrsFreqdensityR16 *[]utils.INTEGER `lb:2,ub:2`
	SlPtrsTimedensityR16 *[]utils.INTEGER `lb:3,ub:3`
	SlPtrsReOffsetR16    *SlPtrsConfigR16SlPtrsReOffsetR16
}

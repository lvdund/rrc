package ies

import "rrc/utils"

// MsgA-DMRS-Config-r16-msgA-MaxLength-r16 ::= ENUMERATED
type MsgaDmrsConfigR16MsgaMaxlengthR16 struct {
	Value utils.ENUMERATED
}

const (
	MsgaDmrsConfigR16MsgaMaxlengthR16EnumeratedNothing = iota
	MsgaDmrsConfigR16MsgaMaxlengthR16EnumeratedLen2
)

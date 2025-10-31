package ies

import "rrc/utils"

// MsgA-DMRS-Config-r16-msgA-DMRS-AdditionalPosition-r16 ::= ENUMERATED
type MsgaDmrsConfigR16MsgaDmrsAdditionalpositionR16 struct {
	Value utils.ENUMERATED
}

const (
	MsgaDmrsConfigR16MsgaDmrsAdditionalpositionR16EnumeratedNothing = iota
	MsgaDmrsConfigR16MsgaDmrsAdditionalpositionR16EnumeratedPos0
	MsgaDmrsConfigR16MsgaDmrsAdditionalpositionR16EnumeratedPos1
	MsgaDmrsConfigR16MsgaDmrsAdditionalpositionR16EnumeratedPos3
)

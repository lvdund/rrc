package ies

import "rrc/utils"

// WidebandPRG-r16 ::= SEQUENCE
type WidebandprgR16 struct {
	WidebandprgSubframeR16    utils.BOOLEAN
	WidebandprgSlotsubslotR16 utils.BOOLEAN
}

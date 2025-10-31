package ies

import "rrc/utils"

// SubSlot-Config-r16-sub-SlotConfig-ECP-r16 ::= ENUMERATED
type SubslotConfigR16SubSlotconfigEcpR16 struct {
	Value utils.ENUMERATED
}

const (
	SubslotConfigR16SubSlotconfigEcpR16EnumeratedNothing = iota
	SubslotConfigR16SubSlotconfigEcpR16EnumeratedN4
	SubslotConfigR16SubSlotconfigEcpR16EnumeratedN5
	SubslotConfigR16SubSlotconfigEcpR16EnumeratedN6
)

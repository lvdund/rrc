package ies

import "rrc/utils"

// SubSlot-Config-r16-sub-SlotConfig-NCP-r16 ::= ENUMERATED
type SubslotConfigR16SubSlotconfigNcpR16 struct {
	Value utils.ENUMERATED
}

const (
	SubslotConfigR16SubSlotconfigNcpR16EnumeratedNothing = iota
	SubslotConfigR16SubSlotconfigNcpR16EnumeratedN4
	SubslotConfigR16SubSlotconfigNcpR16EnumeratedN5
	SubslotConfigR16SubSlotconfigNcpR16EnumeratedN6
	SubslotConfigR16SubSlotconfigNcpR16EnumeratedN7
)

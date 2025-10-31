package ies

import "rrc/utils"

// MsgA-PUSCH-Resource-r16-msgA-IntraSlotFrequencyHopping-r16 ::= ENUMERATED
type MsgaPuschResourceR16MsgaIntraslotfrequencyhoppingR16 struct {
	Value utils.ENUMERATED
}

const (
	MsgaPuschResourceR16MsgaIntraslotfrequencyhoppingR16EnumeratedNothing = iota
	MsgaPuschResourceR16MsgaIntraslotfrequencyhoppingR16EnumeratedEnabled
)

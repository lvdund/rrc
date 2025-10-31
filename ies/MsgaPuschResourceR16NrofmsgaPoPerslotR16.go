package ies

import "rrc/utils"

// MsgA-PUSCH-Resource-r16-nrofMsgA-PO-PerSlot-r16 ::= ENUMERATED
type MsgaPuschResourceR16NrofmsgaPoPerslotR16 struct {
	Value utils.ENUMERATED
}

const (
	MsgaPuschResourceR16NrofmsgaPoPerslotR16EnumeratedNothing = iota
	MsgaPuschResourceR16NrofmsgaPoPerslotR16EnumeratedOne
	MsgaPuschResourceR16NrofmsgaPoPerslotR16EnumeratedTwo
	MsgaPuschResourceR16NrofmsgaPoPerslotR16EnumeratedThree
	MsgaPuschResourceR16NrofmsgaPoPerslotR16EnumeratedSix
)

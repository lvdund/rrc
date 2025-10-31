package ies

import "rrc/utils"

// MsgA-PUSCH-Resource-r16-nrofMsgA-PO-FDM-r16 ::= ENUMERATED
type MsgaPuschResourceR16NrofmsgaPoFdmR16 struct {
	Value utils.ENUMERATED
}

const (
	MsgaPuschResourceR16NrofmsgaPoFdmR16EnumeratedNothing = iota
	MsgaPuschResourceR16NrofmsgaPoFdmR16EnumeratedOne
	MsgaPuschResourceR16NrofmsgaPoFdmR16EnumeratedTwo
	MsgaPuschResourceR16NrofmsgaPoFdmR16EnumeratedFour
	MsgaPuschResourceR16NrofmsgaPoFdmR16EnumeratedEight
)

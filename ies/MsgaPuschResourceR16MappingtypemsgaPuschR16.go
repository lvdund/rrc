package ies

import "rrc/utils"

// MsgA-PUSCH-Resource-r16-mappingTypeMsgA-PUSCH-r16 ::= ENUMERATED
type MsgaPuschResourceR16MappingtypemsgaPuschR16 struct {
	Value utils.ENUMERATED
}

const (
	MsgaPuschResourceR16MappingtypemsgaPuschR16EnumeratedNothing = iota
	MsgaPuschResourceR16MappingtypemsgaPuschR16EnumeratedTypea
	MsgaPuschResourceR16MappingtypemsgaPuschR16EnumeratedTypeb
)

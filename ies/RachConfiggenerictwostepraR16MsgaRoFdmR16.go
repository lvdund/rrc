package ies

import "rrc/utils"

// RACH-ConfigGenericTwoStepRA-r16-msgA-RO-FDM-r16 ::= ENUMERATED
type RachConfiggenerictwostepraR16MsgaRoFdmR16 struct {
	Value utils.ENUMERATED
}

const (
	RachConfiggenerictwostepraR16MsgaRoFdmR16EnumeratedNothing = iota
	RachConfiggenerictwostepraR16MsgaRoFdmR16EnumeratedOne
	RachConfiggenerictwostepraR16MsgaRoFdmR16EnumeratedTwo
	RachConfiggenerictwostepraR16MsgaRoFdmR16EnumeratedFour
	RachConfiggenerictwostepraR16MsgaRoFdmR16EnumeratedEight
)

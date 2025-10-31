package ies

import "rrc/utils"

// UL-AM-RLC-r15 ::= SEQUENCE
type UlAmRlcR15 struct {
	TPollretransmitR15    TPollretransmit
	PollpduR15            PollpduR15
	PollbyteR15           PollbyteR14
	MaxretxthresholdR15   UlAmRlcR15MaxretxthresholdR15
	ExtendedRlcLiFieldR15 utils.BOOLEAN
}

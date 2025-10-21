package ies

import "rrc/utils"

// UL-AM-RLC-NB-r13 ::= SEQUENCE
type UlAmRlcNbR13 struct {
	TPollretransmitR13  TPollretransmitNbR13
	MaxretxthresholdR13 utils.ENUMERATED
}

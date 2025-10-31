package ies

import "rrc/utils"

// P-C-AndCBSR-r11 ::= SEQUENCE
type PCAndcbsrR11 struct {
	PCR11                        utils.INTEGER `lb:0,ub:15`
	CodebooksubsetrestrictionR11 utils.BITSTRING
}

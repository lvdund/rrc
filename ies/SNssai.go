package ies

import "rrc/utils"

// S-NSSAI ::= CHOICE
const (
	SNssaiChoiceNothing = iota
	SNssaiChoiceSst
	SNssaiChoiceSstSd
)

type SNssai struct {
	Choice uint64
	Sst    *utils.BITSTRING `lb:8,ub:8`
	SstSd  *utils.BITSTRING `lb:32,ub:32`
}

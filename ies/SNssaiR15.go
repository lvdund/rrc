package ies

import "rrc/utils"

// S-NSSAI-r15 ::= CHOICE
const (
	SNssaiR15ChoiceNothing = iota
	SNssaiR15ChoiceSst
	SNssaiR15ChoiceSstSd
)

type SNssaiR15 struct {
	Choice uint64
	Sst    *utils.BITSTRING `lb:8,ub:8`
	SstSd  *utils.BITSTRING `lb:32,ub:32`
}

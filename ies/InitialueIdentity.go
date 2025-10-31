package ies

import "rrc/utils"

// InitialUE-Identity ::= CHOICE
const (
	InitialueIdentityChoiceNothing = iota
	InitialueIdentityChoiceNg5gSTmsiPart1
	InitialueIdentityChoiceRandomvalue
)

type InitialueIdentity struct {
	Choice         uint64
	Ng5gSTmsiPart1 *utils.BITSTRING `lb:39,ub:39`
	Randomvalue    *utils.BITSTRING `lb:39,ub:39`
}

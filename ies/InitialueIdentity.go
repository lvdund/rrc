package ies

import "rrc/utils"

// InitialUE-Identity ::= CHOICE
const (
	InitialueIdentityChoiceNothing = iota
	InitialueIdentityChoiceSTmsi
	InitialueIdentityChoiceRandomvalue
)

type InitialueIdentity struct {
	Choice      uint64
	STmsi       *STmsi
	Randomvalue *utils.BITSTRING `lb:40,ub:40`
}

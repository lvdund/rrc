package ies

import "rrc/utils"

// InitialUE-Identity-5GC-NB-r16 ::= CHOICE
const (
	InitialueIdentity5gcNbR16ChoiceNothing = iota
	InitialueIdentity5gcNbR16ChoiceNg5gSTmsiR16
	InitialueIdentity5gcNbR16ChoiceRandomvalue
)

type InitialueIdentity5gcNbR16 struct {
	Choice       uint64
	Ng5gSTmsiR16 *Ng5gSTmsiR15
	Randomvalue  *utils.BITSTRING `lb:48,ub:48`
}

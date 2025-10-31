package ies

import "rrc/utils"

// InitialUE-Identity-5GC-r15 ::= CHOICE
const (
	InitialueIdentity5gcR15ChoiceNothing = iota
	InitialueIdentity5gcR15ChoiceNg5gSTmsiPart1
	InitialueIdentity5gcR15ChoiceRandomvalue
)

type InitialueIdentity5gcR15 struct {
	Choice         uint64
	Ng5gSTmsiPart1 *utils.BITSTRING `lb:40,ub:40`
	Randomvalue    *utils.BITSTRING `lb:40,ub:40`
}

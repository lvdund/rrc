package ies

import "rrc/utils"

// InitialUE-Identity-5GC-NB-r16 ::= CHOICE
type InitialueIdentity5gcNbR16 interface {
	isInitialueIdentity5gcNbR16()
}

type InitialueIdentity5gcNbR16Ng5gSTmsiR16 struct {
	Value Ng5gSTmsiR15
}

func (*InitialueIdentity5gcNbR16Ng5gSTmsiR16) isInitialueIdentity5gcNbR16() {}

type InitialueIdentity5gcNbR16Randomvalue struct {
	Value utils.BITSTRING
}

func (*InitialueIdentity5gcNbR16Randomvalue) isInitialueIdentity5gcNbR16() {}

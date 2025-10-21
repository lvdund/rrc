package ies

import "rrc/utils"

// InitialUE-Identity-5GC-r15 ::= CHOICE
type InitialueIdentity5gcR15 interface {
	isInitialueIdentity5gcR15()
}

type InitialueIdentity5gcR15Ng5gSTmsiPart1 struct {
	Value utils.BITSTRING
}

func (*InitialueIdentity5gcR15Ng5gSTmsiPart1) isInitialueIdentity5gcR15() {}

type InitialueIdentity5gcR15Randomvalue struct {
	Value utils.BITSTRING
}

func (*InitialueIdentity5gcR15Randomvalue) isInitialueIdentity5gcR15() {}

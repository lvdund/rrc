package ies

import "rrc/utils"

// InitialUE-Identity ::= CHOICE
type InitialueIdentity interface {
	isInitialueIdentity()
}

type InitialueIdentitySTmsi struct {
	Value STmsi
}

func (*InitialueIdentitySTmsi) isInitialueIdentity() {}

type InitialueIdentityRandomvalue struct {
	Value utils.BITSTRING
}

func (*InitialueIdentityRandomvalue) isInitialueIdentity() {}

package ies

import "rrc/utils"

// PagingUE-Identity ::= CHOICE
// Extensible
type PagingueIdentity interface {
	isPagingueIdentity()
}

type PagingueIdentitySTmsi struct {
	Value STmsi
}

func (*PagingueIdentitySTmsi) isPagingueIdentity() {}

type PagingueIdentityImsi struct {
	Value Imsi
}

func (*PagingueIdentityImsi) isPagingueIdentity() {}

type PagingueIdentityNg5gSTmsiR15 struct {
	Value Ng5gSTmsiR15
}

func (*PagingueIdentityNg5gSTmsiR15) isPagingueIdentity() {}

type PagingueIdentityFulliRntiR15 struct {
	Value IRntiR15
}

func (*PagingueIdentityFulliRntiR15) isPagingueIdentity() {}

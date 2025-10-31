package ies

// PagingUE-Identity ::= CHOICE
// Extensible
const (
	PagingueIdentityChoiceNothing = iota
	PagingueIdentityChoiceSTmsi
	PagingueIdentityChoiceImsi
	PagingueIdentityChoiceNg5gSTmsiR15
	PagingueIdentityChoiceFulliRntiR15
)

type PagingueIdentity struct {
	Choice       uint64
	STmsi        *STmsi
	Imsi         *Imsi
	Ng5gSTmsiR15 *Ng5gSTmsiR15
	FulliRntiR15 *IRntiR15
}

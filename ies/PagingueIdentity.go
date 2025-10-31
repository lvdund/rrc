package ies

// PagingUE-Identity ::= CHOICE
// Extensible
const (
	PagingueIdentityChoiceNothing = iota
	PagingueIdentityChoiceNg5gSTmsi
	PagingueIdentityChoiceFulliRnti
)

type PagingueIdentity struct {
	Choice    uint64
	Ng5gSTmsi *Ng5gSTmsi
	FulliRnti *IRntiValue
}

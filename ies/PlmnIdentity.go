package ies

// PLMN-Identity ::= SEQUENCE
type PlmnIdentity struct {
	Mcc *Mcc
	Mnc Mnc
}

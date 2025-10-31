package ies

// PLMN-IdentityList-5GC-NB-r16 ::= SEQUENCE OF PLMN-IdentityInfo-5GC-NB-r16
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylist5gcNbR16 struct {
	Value []PlmnIdentityinfo5gcNbR16 `lb:1,ub:maxPLMNR11`
}

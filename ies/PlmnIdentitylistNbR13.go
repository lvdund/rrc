package ies

// PLMN-IdentityList-NB-r13 ::= SEQUENCE OF PLMN-IdentityInfo-NB-r13
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylistNbR13 struct {
	Value []PlmnIdentityinfoNbR13 `lb:1,ub:maxPLMNR11`
}

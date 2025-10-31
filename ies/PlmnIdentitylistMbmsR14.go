package ies

// PLMN-IdentityList-MBMS-r14 ::= SEQUENCE OF PLMN-Identity
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylistMbmsR14 struct {
	Value []PlmnIdentity `lb:1,ub:maxPLMNR11`
}

package ies

// PLMN-IdentityList ::= SEQUENCE OF PLMN-IdentityInfo
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylist struct {
	Value []PlmnIdentityinfo `lb:1,ub:maxPLMNR11`
}

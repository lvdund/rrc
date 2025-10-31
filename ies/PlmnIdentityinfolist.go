package ies

// PLMN-IdentityInfoList ::= SEQUENCE OF PLMN-IdentityInfo
// SIZE (1..maxPLMN)
type PlmnIdentityinfolist struct {
	Value []PlmnIdentityinfo `lb:1,ub:maxPLMN`
}

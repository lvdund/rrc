package ies

// PLMN-IdentityList-r16 ::= SEQUENCE OF PLMN-Identity
// SIZE (1..maxPLMN)
type PlmnIdentitylistR16 struct {
	Value []PlmnIdentity `lb:1,ub:maxPLMN`
}

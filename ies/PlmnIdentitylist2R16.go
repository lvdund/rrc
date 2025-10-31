package ies

// PLMN-IdentityList2-r16 ::= SEQUENCE OF PLMN-Identity
// SIZE (1..16)
type PlmnIdentitylist2R16 struct {
	Value []PlmnIdentity `lb:1,ub:16`
}

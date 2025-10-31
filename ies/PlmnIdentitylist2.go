package ies

// PLMN-IdentityList2 ::= SEQUENCE OF PLMN-Identity
// SIZE (1..5)
type PlmnIdentitylist2 struct {
	Value []PlmnIdentity `lb:1,ub:5`
}

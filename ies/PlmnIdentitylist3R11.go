package ies

// PLMN-IdentityList3-r11 ::= SEQUENCE OF PLMN-Identity
// SIZE (1..16)
type PlmnIdentitylist3R11 struct {
	Value []PlmnIdentity `lb:1,ub:16`
}

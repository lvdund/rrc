package ies

// PLMN-IdentityListNR-r15 ::= SEQUENCE OF PLMN-Identity
// SIZE (1.. maxPLMN-NR-r15)
type PlmnIdentitylistnrR15 struct {
	Value []PlmnIdentity `lb:1,ub:maxPLMNNrR15`
}

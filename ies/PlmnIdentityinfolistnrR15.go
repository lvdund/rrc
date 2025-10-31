package ies

// PLMN-IdentityInfoListNR-r15 ::= SEQUENCE OF PLMN-IdentityInfoNR-r15
// SIZE (1..maxPLMN-NR-r15)
type PlmnIdentityinfolistnrR15 struct {
	Value []PlmnIdentityinfonrR15 `lb:1,ub:maxPLMNNrR15`
}

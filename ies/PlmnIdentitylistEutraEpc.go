package ies

// PLMN-IdentityList-EUTRA-EPC ::= SEQUENCE OF PLMN-Identity
// SIZE (1..maxPLMN)
type PlmnIdentitylistEutraEpc struct {
	Value []PlmnIdentity `lb:1,ub:maxPLMN`
}

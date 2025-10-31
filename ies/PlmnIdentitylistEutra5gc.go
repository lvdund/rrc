package ies

// PLMN-IdentityList-EUTRA-5GC ::= SEQUENCE OF PLMN-Identity-EUTRA-5GC
// SIZE (1..maxPLMN)
type PlmnIdentitylistEutra5gc struct {
	Value []PlmnIdentityEutra5gc `lb:1,ub:maxPLMN`
}

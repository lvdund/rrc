package ies

// MRB-ToReleaseList-r17 ::= SEQUENCE OF MRB-Identity-r17
// SIZE (1..maxMRB-r17)
type MrbToreleaselistR17 struct {
	Value []MrbIdentityR17 `lb:1,ub:maxMRBR17`
}

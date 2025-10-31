package ies

// DRB-ToReleaseList-r15 ::= SEQUENCE OF DRB-Identity
// SIZE (1..maxDRB-r15)
type DrbToreleaselistR15 struct {
	Value []DrbIdentity `lb:1,ub:maxDRBR15`
}

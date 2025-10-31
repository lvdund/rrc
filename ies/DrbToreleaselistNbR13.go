package ies

// DRB-ToReleaseList-NB-r13 ::= SEQUENCE OF DRB-Identity
// SIZE (1..maxDRB-NB-r13)
type DrbToreleaselistNbR13 struct {
	Value []DrbIdentity `lb:1,ub:maxDRBNbR13`
}

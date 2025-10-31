package ies

// DRB-ToReleaseList ::= SEQUENCE OF DRB-Identity
// SIZE (1..maxDRB)
type DrbToreleaselist struct {
	Value []DrbIdentity `lb:1,ub:maxDRB`
}

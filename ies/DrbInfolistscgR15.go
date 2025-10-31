package ies

// DRB-InfoListSCG-r15 ::= SEQUENCE OF DRB-InfoSCG-r12
// SIZE (1..maxDRB-r15)
type DrbInfolistscgR15 struct {
	Value []DrbInfoscgR12 `lb:1,ub:maxDRBR15`
}

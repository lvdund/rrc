package ies

// DRB-ToAddModListSCG-r15 ::= SEQUENCE OF DRB-ToAddModSCG-r12
// SIZE (1..maxDRB-r15)
type DrbToaddmodlistscgR15 struct {
	Value []DrbToaddmodscgR12 `lb:1,ub:maxDRBR15`
}

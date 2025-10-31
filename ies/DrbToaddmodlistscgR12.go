package ies

// DRB-ToAddModListSCG-r12 ::= SEQUENCE OF DRB-ToAddModSCG-r12
// SIZE (1..maxDRB)
type DrbToaddmodlistscgR12 struct {
	Value []DrbToaddmodscgR12 `lb:1,ub:maxDRB`
}

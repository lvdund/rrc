package ies

// MRB-ToAddModList-r17 ::= SEQUENCE OF MRB-ToAddMod-r17
// SIZE (1..maxMRB-r17)
type MrbToaddmodlistR17 struct {
	Value []MrbToaddmodR17 `lb:1,ub:maxMRBR17`
}

package ies

// SRB-ToAddModList-NB-r13 ::= SEQUENCE OF SRB-ToAddMod-NB-r13
// SIZE (1)
type SrbToaddmodlistNbR13 struct {
	Value []SrbToaddmodNbR13 `lb:1,ub:1`
}

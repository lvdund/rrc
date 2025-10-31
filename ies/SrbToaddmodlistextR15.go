package ies

// SRB-ToAddModListExt-r15 ::= SEQUENCE OF SRB-ToAddMod
// SIZE (1)
type SrbToaddmodlistextR15 struct {
	Value []SrbToaddmod `lb:1,ub:1`
}

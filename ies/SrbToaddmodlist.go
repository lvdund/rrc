package ies

// SRB-ToAddModList ::= SEQUENCE OF SRB-ToAddMod
// SIZE (1..2)
type SrbToaddmodlist struct {
	Value []SrbToaddmod `lb:1,ub:2`
}

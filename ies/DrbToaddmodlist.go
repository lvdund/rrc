package ies

// DRB-ToAddModList ::= SEQUENCE OF DRB-ToAddMod
// SIZE (1..maxDRB)
type DrbToaddmodlist struct {
	Value []DrbToaddmod `lb:1,ub:maxDRB`
}

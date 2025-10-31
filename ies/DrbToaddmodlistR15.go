package ies

// DRB-ToAddModList-r15 ::= SEQUENCE OF DRB-ToAddMod
// SIZE (1..maxDRB-r15)
type DrbToaddmodlistR15 struct {
	Value []DrbToaddmod `lb:1,ub:maxDRBR15`
}

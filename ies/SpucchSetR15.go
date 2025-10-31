package ies

// SPUCCH-Set-r15 ::= SEQUENCE OF SPUCCH-Elements-r15
// SIZE (1..4)
type SpucchSetR15 struct {
	Value []SpucchElementsR15 `lb:1,ub:4`
}

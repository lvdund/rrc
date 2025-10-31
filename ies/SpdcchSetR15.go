package ies

// SPDCCH-Set-r15 ::= SEQUENCE OF SPDCCH-Elements-r15
// SIZE (1..4)
type SpdcchSetR15 struct {
	Value []SpdcchElementsR15 `lb:1,ub:4`
}

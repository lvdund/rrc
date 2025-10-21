package ies

import "rrc/utils"

// SPUCCH-Set-r15 ::= SEQUENCE OF SPUCCH-Elements-r15
// SIZE (1..4)
type SpucchSetR15 struct {
	Value []SpucchElementsR15
}

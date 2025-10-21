package ies

import "rrc/utils"

// STTI-SupportedCombinations-r15 ::= SEQUENCE
type SttiSupportedcombinationsR15 struct {
	Combination22R15   *DlUlCcsR15
	Combination77R15   *DlUlCcsR15
	Combination27R15   *DlUlCcsR15
	Combination2227R15 *interface{}
	Combination7722R15 *interface{}
	Combination7727R15 *interface{}
}

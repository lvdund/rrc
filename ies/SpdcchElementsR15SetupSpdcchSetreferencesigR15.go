package ies

import "rrc/utils"

// SPDCCH-Elements-r15-setup-spdcch-SetReferenceSig-r15 ::= ENUMERATED
type SpdcchElementsR15SetupSpdcchSetreferencesigR15 struct {
	Value utils.ENUMERATED
}

const (
	SpdcchElementsR15SetupSpdcchSetreferencesigR15EnumeratedNothing = iota
	SpdcchElementsR15SetupSpdcchSetreferencesigR15EnumeratedCrs
	SpdcchElementsR15SetupSpdcchSetreferencesigR15EnumeratedDmrs
)

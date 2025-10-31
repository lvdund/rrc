package ies

import "rrc/utils"

// SPDCCH-Elements-r15-setup-transmissionType-r15 ::= ENUMERATED
type SpdcchElementsR15SetupTransmissiontypeR15 struct {
	Value utils.ENUMERATED
}

const (
	SpdcchElementsR15SetupTransmissiontypeR15EnumeratedNothing = iota
	SpdcchElementsR15SetupTransmissiontypeR15EnumeratedLocalised
	SpdcchElementsR15SetupTransmissiontypeR15EnumeratedDistributed
)

package ies

import "rrc/utils"

// SPDCCH-Elements-r15-setup-subframeType-r15 ::= ENUMERATED
type SpdcchElementsR15SetupSubframetypeR15 struct {
	Value utils.ENUMERATED
}

const (
	SpdcchElementsR15SetupSubframetypeR15EnumeratedNothing = iota
	SpdcchElementsR15SetupSubframetypeR15EnumeratedMbsfn
	SpdcchElementsR15SetupSubframetypeR15EnumeratedNonmbsfn
	SpdcchElementsR15SetupSubframetypeR15EnumeratedAll
)

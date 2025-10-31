package ies

import "rrc/utils"

// SPDCCH-Elements-r15-setup-rateMatchingMode-r15 ::= ENUMERATED
type SpdcchElementsR15SetupRatematchingmodeR15 struct {
	Value utils.ENUMERATED
}

const (
	SpdcchElementsR15SetupRatematchingmodeR15EnumeratedNothing = iota
	SpdcchElementsR15SetupRatematchingmodeR15EnumeratedM1
	SpdcchElementsR15SetupRatematchingmodeR15EnumeratedM2
	SpdcchElementsR15SetupRatematchingmodeR15EnumeratedM3
	SpdcchElementsR15SetupRatematchingmodeR15EnumeratedM4
)

package ies

import "rrc/utils"

// WUS-Config-NB-r15-timeOffset-eDRX-Short-r15 ::= ENUMERATED
type WusConfigNbR15TimeoffsetEdrxShortR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigNbR15TimeoffsetEdrxShortR15EnumeratedNothing = iota
	WusConfigNbR15TimeoffsetEdrxShortR15EnumeratedMs40
	WusConfigNbR15TimeoffsetEdrxShortR15EnumeratedMs80
	WusConfigNbR15TimeoffsetEdrxShortR15EnumeratedMs160
	WusConfigNbR15TimeoffsetEdrxShortR15EnumeratedMs240
)

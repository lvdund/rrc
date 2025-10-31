package ies

import "rrc/utils"

// WUS-Config-r15-timeOffset-eDRX-Short-r15 ::= ENUMERATED
type WusConfigR15TimeoffsetEdrxShortR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigR15TimeoffsetEdrxShortR15EnumeratedNothing = iota
	WusConfigR15TimeoffsetEdrxShortR15EnumeratedMs40
	WusConfigR15TimeoffsetEdrxShortR15EnumeratedMs80
	WusConfigR15TimeoffsetEdrxShortR15EnumeratedMs160
	WusConfigR15TimeoffsetEdrxShortR15EnumeratedMs240
)

package ies

import "rrc/utils"

// WUS-Config-r15-timeOffset-eDRX-Long-r15 ::= ENUMERATED
type WusConfigR15TimeoffsetEdrxLongR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigR15TimeoffsetEdrxLongR15EnumeratedNothing = iota
	WusConfigR15TimeoffsetEdrxLongR15EnumeratedMs1000
	WusConfigR15TimeoffsetEdrxLongR15EnumeratedMs2000
)

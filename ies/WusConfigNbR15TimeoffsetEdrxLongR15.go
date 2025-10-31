package ies

import "rrc/utils"

// WUS-Config-NB-r15-timeOffset-eDRX-Long-r15 ::= ENUMERATED
type WusConfigNbR15TimeoffsetEdrxLongR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigNbR15TimeoffsetEdrxLongR15EnumeratedNothing = iota
	WusConfigNbR15TimeoffsetEdrxLongR15EnumeratedMs1000
	WusConfigNbR15TimeoffsetEdrxLongR15EnumeratedMs2000
)

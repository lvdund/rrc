package ies

import "rrc/utils"

// WUS-Config-NB-r15-timeOffsetDRX-r15 ::= ENUMERATED
type WusConfigNbR15TimeoffsetdrxR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigNbR15TimeoffsetdrxR15EnumeratedNothing = iota
	WusConfigNbR15TimeoffsetdrxR15EnumeratedMs40
	WusConfigNbR15TimeoffsetdrxR15EnumeratedMs80
	WusConfigNbR15TimeoffsetdrxR15EnumeratedMs160
	WusConfigNbR15TimeoffsetdrxR15EnumeratedMs240
)

package ies

import "rrc/utils"

// WUS-Config-r15-timeOffsetDRX-r15 ::= ENUMERATED
type WusConfigR15TimeoffsetdrxR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigR15TimeoffsetdrxR15EnumeratedNothing = iota
	WusConfigR15TimeoffsetdrxR15EnumeratedMs40
	WusConfigR15TimeoffsetdrxR15EnumeratedMs80
	WusConfigR15TimeoffsetdrxR15EnumeratedMs160
	WusConfigR15TimeoffsetdrxR15EnumeratedMs240
)

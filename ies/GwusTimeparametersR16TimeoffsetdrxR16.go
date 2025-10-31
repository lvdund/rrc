package ies

import "rrc/utils"

// GWUS-TimeParameters-r16-timeOffsetDRX-r16 ::= ENUMERATED
type GwusTimeparametersR16TimeoffsetdrxR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusTimeparametersR16TimeoffsetdrxR16EnumeratedNothing = iota
	GwusTimeparametersR16TimeoffsetdrxR16EnumeratedMs40
	GwusTimeparametersR16TimeoffsetdrxR16EnumeratedMs80
	GwusTimeparametersR16TimeoffsetdrxR16EnumeratedMs160
	GwusTimeparametersR16TimeoffsetdrxR16EnumeratedMs240
)

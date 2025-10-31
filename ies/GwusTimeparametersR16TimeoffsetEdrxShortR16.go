package ies

import "rrc/utils"

// GWUS-TimeParameters-r16-timeOffset-eDRX-Short-r16 ::= ENUMERATED
type GwusTimeparametersR16TimeoffsetEdrxShortR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusTimeparametersR16TimeoffsetEdrxShortR16EnumeratedNothing = iota
	GwusTimeparametersR16TimeoffsetEdrxShortR16EnumeratedMs40
	GwusTimeparametersR16TimeoffsetEdrxShortR16EnumeratedMs80
	GwusTimeparametersR16TimeoffsetEdrxShortR16EnumeratedMs160
	GwusTimeparametersR16TimeoffsetEdrxShortR16EnumeratedMs240
)

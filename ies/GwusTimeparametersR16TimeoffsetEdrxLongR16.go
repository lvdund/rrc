package ies

import "rrc/utils"

// GWUS-TimeParameters-r16-timeOffset-eDRX-Long-r16 ::= ENUMERATED
type GwusTimeparametersR16TimeoffsetEdrxLongR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusTimeparametersR16TimeoffsetEdrxLongR16EnumeratedNothing = iota
	GwusTimeparametersR16TimeoffsetEdrxLongR16EnumeratedMs1000
	GwusTimeparametersR16TimeoffsetEdrxLongR16EnumeratedMs2000
)

package ies

import "rrc/utils"

// GWUS-TimeParameters-r16-numDRX-CyclesRelaxed-r16 ::= ENUMERATED
type GwusTimeparametersR16NumdrxCyclesrelaxedR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusTimeparametersR16NumdrxCyclesrelaxedR16EnumeratedNothing = iota
	GwusTimeparametersR16NumdrxCyclesrelaxedR16EnumeratedN1
	GwusTimeparametersR16NumdrxCyclesrelaxedR16EnumeratedN2
	GwusTimeparametersR16NumdrxCyclesrelaxedR16EnumeratedN4
	GwusTimeparametersR16NumdrxCyclesrelaxedR16EnumeratedN8
)

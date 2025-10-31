package ies

import "rrc/utils"

// WUS-Config-v1610-numDRX-CyclesRelaxed-r16 ::= ENUMERATED
type WusConfigV1610NumdrxCyclesrelaxedR16 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigV1610NumdrxCyclesrelaxedR16EnumeratedNothing = iota
	WusConfigV1610NumdrxCyclesrelaxedR16EnumeratedN1
	WusConfigV1610NumdrxCyclesrelaxedR16EnumeratedN2
	WusConfigV1610NumdrxCyclesrelaxedR16EnumeratedN4
	WusConfigV1610NumdrxCyclesrelaxedR16EnumeratedN8
)

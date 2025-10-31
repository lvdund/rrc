package ies

import "rrc/utils"

// WUS-Config-NB-r15-numDRX-CyclesRelaxed-r15 ::= ENUMERATED
type WusConfigNbR15NumdrxCyclesrelaxedR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigNbR15NumdrxCyclesrelaxedR15EnumeratedNothing = iota
	WusConfigNbR15NumdrxCyclesrelaxedR15EnumeratedN1
	WusConfigNbR15NumdrxCyclesrelaxedR15EnumeratedN2
	WusConfigNbR15NumdrxCyclesrelaxedR15EnumeratedN4
	WusConfigNbR15NumdrxCyclesrelaxedR15EnumeratedN8
)

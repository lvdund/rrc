package ies

import "rrc/utils"

// PUSCH-PowerControl-tpc-Accumulation ::= ENUMERATED
type PuschPowercontrolTpcAccumulation struct {
	Value utils.ENUMERATED
}

const (
	PuschPowercontrolTpcAccumulationEnumeratedNothing = iota
	PuschPowercontrolTpcAccumulationEnumeratedDisabled
)

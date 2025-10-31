package ies

import "rrc/utils"

// PUSCH-PowerControl-twoPUSCH-PC-AdjustmentStates ::= ENUMERATED
type PuschPowercontrolTwopuschPcAdjustmentstates struct {
	Value utils.ENUMERATED
}

const (
	PuschPowercontrolTwopuschPcAdjustmentstatesEnumeratedNothing = iota
	PuschPowercontrolTwopuschPcAdjustmentstatesEnumeratedTwostates
)

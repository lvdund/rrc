package ies

import "rrc/utils"

// PUCCH-PowerControl-twoPUCCH-PC-AdjustmentStates ::= ENUMERATED
type PucchPowercontrolTwopucchPcAdjustmentstates struct {
	Value utils.ENUMERATED
}

const (
	PucchPowercontrolTwopucchPcAdjustmentstatesEnumeratedNothing = iota
	PucchPowercontrolTwopucchPcAdjustmentstatesEnumeratedTwostates
)

package ies

import "rrc/utils"

// STTI-SPT-BandParameters-r15-ul-256QAM-Slot-r15 ::= ENUMERATED
type SttiSptBandparametersR15Ul256qamSlotR15 struct {
	Value utils.ENUMERATED
}

const (
	SttiSptBandparametersR15Ul256qamSlotR15EnumeratedNothing = iota
	SttiSptBandparametersR15Ul256qamSlotR15EnumeratedSupported
)

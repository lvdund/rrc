package ies

import "rrc/utils"

// STTI-SPT-BandParameters-r15-dl-1024QAM-Slot-r15 ::= ENUMERATED
type SttiSptBandparametersR15Dl1024qamSlotR15 struct {
	Value utils.ENUMERATED
}

const (
	SttiSptBandparametersR15Dl1024qamSlotR15EnumeratedNothing = iota
	SttiSptBandparametersR15Dl1024qamSlotR15EnumeratedSupported
)

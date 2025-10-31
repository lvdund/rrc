package ies

import "rrc/utils"

// STTI-SPT-BandParameters-r15-ul-256QAM-Subslot-r15 ::= ENUMERATED
type SttiSptBandparametersR15Ul256qamSubslotR15 struct {
	Value utils.ENUMERATED
}

const (
	SttiSptBandparametersR15Ul256qamSubslotR15EnumeratedNothing = iota
	SttiSptBandparametersR15Ul256qamSubslotR15EnumeratedSupported
)

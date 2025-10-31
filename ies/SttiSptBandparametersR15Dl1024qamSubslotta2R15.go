package ies

import "rrc/utils"

// STTI-SPT-BandParameters-r15-dl-1024QAM-SubslotTA-2-r15 ::= ENUMERATED
type SttiSptBandparametersR15Dl1024qamSubslotta2R15 struct {
	Value utils.ENUMERATED
}

const (
	SttiSptBandparametersR15Dl1024qamSubslotta2R15EnumeratedNothing = iota
	SttiSptBandparametersR15Dl1024qamSubslotta2R15EnumeratedSupported
)

package ies

import "rrc/utils"

// STTI-SPT-BandParameters-r15-simultaneousTx-differentTx-duration-r15 ::= ENUMERATED
type SttiSptBandparametersR15SimultaneoustxDifferenttxDurationR15 struct {
	Value utils.ENUMERATED
}

const (
	SttiSptBandparametersR15SimultaneoustxDifferenttxDurationR15EnumeratedNothing = iota
	SttiSptBandparametersR15SimultaneoustxDifferenttxDurationR15EnumeratedSupported
)

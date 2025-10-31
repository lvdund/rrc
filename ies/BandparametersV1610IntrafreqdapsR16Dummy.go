package ies

import "rrc/utils"

// BandParameters-v1610-intraFreqDAPS-r16-dummy ::= ENUMERATED
type BandparametersV1610IntrafreqdapsR16Dummy struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1610IntrafreqdapsR16DummyEnumeratedNothing = iota
	BandparametersV1610IntrafreqdapsR16DummyEnumeratedSupported
)

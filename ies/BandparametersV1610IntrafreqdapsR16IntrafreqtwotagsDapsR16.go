package ies

import "rrc/utils"

// BandParameters-v1610-intraFreqDAPS-r16-intraFreqTwoTAGs-DAPS-r16 ::= ENUMERATED
type BandparametersV1610IntrafreqdapsR16IntrafreqtwotagsDapsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1610IntrafreqdapsR16IntrafreqtwotagsDapsR16EnumeratedNothing = iota
	BandparametersV1610IntrafreqdapsR16IntrafreqtwotagsDapsR16EnumeratedSupported
)

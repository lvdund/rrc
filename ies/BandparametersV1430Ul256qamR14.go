package ies

import "rrc/utils"

// BandParameters-v1430-ul-256QAM-r14 ::= ENUMERATED
type BandparametersV1430Ul256qamR14 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1430Ul256qamR14EnumeratedNothing = iota
	BandparametersV1430Ul256qamR14EnumeratedSupported
)

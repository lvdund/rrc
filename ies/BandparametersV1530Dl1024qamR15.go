package ies

import "rrc/utils"

// BandParameters-v1530-dl-1024QAM-r15 ::= ENUMERATED
type BandparametersV1530Dl1024qamR15 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersV1530Dl1024qamR15EnumeratedNothing = iota
	BandparametersV1530Dl1024qamR15EnumeratedSupported
)

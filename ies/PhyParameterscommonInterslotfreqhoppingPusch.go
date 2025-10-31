package ies

import "rrc/utils"

// Phy-ParametersCommon-interSlotFreqHopping-PUSCH ::= ENUMERATED
type PhyParameterscommonInterslotfreqhoppingPusch struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonInterslotfreqhoppingPuschEnumeratedNothing = iota
	PhyParameterscommonInterslotfreqhoppingPuschEnumeratedSupported
)

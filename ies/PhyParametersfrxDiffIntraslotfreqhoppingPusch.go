package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-intraSlotFreqHopping-PUSCH ::= ENUMERATED
type PhyParametersfrxDiffIntraslotfreqhoppingPusch struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffIntraslotfreqhoppingPuschEnumeratedNothing = iota
	PhyParametersfrxDiffIntraslotfreqhoppingPuschEnumeratedSupported
)

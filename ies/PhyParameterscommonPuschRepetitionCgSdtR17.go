package ies

import "rrc/utils"

// Phy-ParametersCommon-pusch-Repetition-CG-SDT-r17 ::= ENUMERATED
type PhyParameterscommonPuschRepetitionCgSdtR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPuschRepetitionCgSdtR17EnumeratedNothing = iota
	PhyParameterscommonPuschRepetitionCgSdtR17EnumeratedSupported
)

package ies

import "rrc/utils"

// Phy-ParametersCommon-pucch-Repetition-F1-3-4 ::= ENUMERATED
type PhyParameterscommonPucchRepetitionF134 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPucchRepetitionF134EnumeratedNothing = iota
	PhyParameterscommonPucchRepetitionF134EnumeratedSupported
)

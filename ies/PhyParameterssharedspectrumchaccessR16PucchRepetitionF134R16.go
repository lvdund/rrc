package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-pucch-Repetition-F1-3-4-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16PucchRepetitionF134R16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16PucchRepetitionF134R16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16PucchRepetitionF134R16EnumeratedSupported
)

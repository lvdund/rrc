package ies

import "rrc/utils"

// Phy-ParametersCommon-pusch-RepetitionTypeA-r16-sharedSpectrumChAccess-r16 ::= ENUMERATED
type PhyParameterscommonPuschRepetitiontypeaR16SharedspectrumchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPuschRepetitiontypeaR16SharedspectrumchaccessR16EnumeratedNothing = iota
	PhyParameterscommonPuschRepetitiontypeaR16SharedspectrumchaccessR16EnumeratedSupported
)

package ies

import "rrc/utils"

// Phy-ParametersCommon-pusch-RepetitionTypeA-r16-non-SharedSpectrumChAccess-r16 ::= ENUMERATED
type PhyParameterscommonPuschRepetitiontypeaR16NonSharedspectrumchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPuschRepetitiontypeaR16NonSharedspectrumchaccessR16EnumeratedNothing = iota
	PhyParameterscommonPuschRepetitiontypeaR16NonSharedspectrumchaccessR16EnumeratedSupported
)

package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-configuredUL-GrantType2-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16ConfiguredulGranttype2R16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16ConfiguredulGranttype2R16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16ConfiguredulGranttype2R16EnumeratedSupported
)

package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-configuredUL-GrantType1-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16ConfiguredulGranttype1R16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16ConfiguredulGranttype1R16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16ConfiguredulGranttype1R16EnumeratedSupported
)

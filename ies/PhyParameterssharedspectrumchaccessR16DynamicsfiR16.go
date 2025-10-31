package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-dynamicSFI-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16DynamicsfiR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16DynamicsfiR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16DynamicsfiR16EnumeratedSupported
)
